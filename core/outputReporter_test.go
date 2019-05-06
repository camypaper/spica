package core_test

import (
	"testing"

	"github.com/Camypaper/libra"
	"github.com/Camypaper/spica/core"
)

type fake struct {
	name string
	libra.Job
	libra.Task
}

func (f fake) Name() string {
	return f.name
}
func TestOutputReporter1(t *testing.T) {
	reporter := core.OutputReporter{}
	f := fake{name: "test"}
	reporter.ReportStart(f)
	reporter.Report(f, libra.Status{Code: libra.OK, Msg: "ok"})
	reporter.Report(f, libra.Status{Code: libra.RE, Msg: "fail"})
	reporter.ReportEnd()
	res := reporter.Get()
	if len(res) != 1 {
		t.Errorf("expected: %v, actual: %v", 1, len(res))
	}
	for _, v := range res {
		if v.Msg != "ok" {
			t.Errorf("expected: %v, actual: %v", libra.Status{Code: libra.OK, Msg: "ok"}, v)
		}
	}
}

func TestOutputReporter2(t *testing.T) {
	reporter := core.OutputReporter{}
	reporter.All = true
	f := fake{name: "test"}
	reporter.ReportStart(f)
	f.name = "test1"
	reporter.Report(f, libra.Status{Code: libra.OK, Msg: "ok"})
	f.name = "test2"
	reporter.Report(f, libra.Status{Code: libra.RE, Msg: "fail"})
	reporter.ReportEnd()
	res := reporter.Get()
	if len(res) != 2 {
		t.Errorf("expected: %v, actual: %v", 1, len(res))
	}
	for k, v := range res {
		if k == "test1" && v.Msg != "ok" {
			t.Errorf("expected: %v, actual: %v", libra.Status{Code: libra.OK, Msg: "ok"}, v)
		}
		if k == "test2" && v.Msg != "fail" {
			t.Errorf("expected: %v, actual: %v", libra.Status{Code: libra.RE, Msg: "fail"}, v)
		}
	}
}
