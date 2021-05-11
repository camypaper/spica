package core

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/camypaper/libra"

	"github.com/sirupsen/logrus"
	pb "gopkg.in/cheggaaa/pb.v2"
)

/*
StdoutReporter shows progress of job.
*/
type StdoutReporter struct {
	success int32
	fail    int32
	all     int
	target  string
	bar     *pb.ProgressBar
	dic     sync.Map
}

/*
ReportStart initialize progress bar.
*/
func (reporter *StdoutReporter) ReportStart(job libra.Job) {
	reporter.success = 0
	reporter.fail = 0
	reporter.all = len(job.Subtasks())
	reporter.bar = pb.StartNew(reporter.all)
	reporter.target = job.Name()
	tmpl := fmt.Sprintf(`%-12s: {{counters . }} {{bar . "[" "#" "#" " " "]" | green}} `, reporter.target)
	reporter.bar.SetTemplateString(tmpl)
	reporter.dic = sync.Map{}
}

/*
Report updates progress.
*/
func (reporter *StdoutReporter) Report(task libra.Task, status libra.Status) {
	reporter.bar.Increment()
	if status.Code == libra.OK {
		atomic.AddInt32(&reporter.success, 1)
	} else {
		if atomic.AddInt32(&reporter.fail, 1) == 1 {
			tmpl := fmt.Sprintf(`%-20s: {{counters . }} {{bar . "[" "#" "#" " " "]" | red}} `, reporter.target)
			reporter.bar.SetTemplateString(tmpl)
		}
		reporter.dic.LoadOrStore(task.Name(), status)
	}
}

/*
ReportEnd shows result.
*/
func (reporter *StdoutReporter) ReportEnd() {
	if reporter.success == int32(reporter.all) {
		tmpl := fmt.Sprintf(`%-20s: {{green "OK"}} `, reporter.target)
		reporter.bar.SetTemplateString(tmpl)

		reporter.bar.Finish()
	} else {
		tmpl := fmt.Sprintf(`%-20s: {{red "NG"}} `, reporter.target)
		reporter.bar.SetTemplateString(tmpl)
		reporter.bar.Finish()

		res := logrus.Fields{}
		reporter.dic.Range(func(key, value interface{}) bool {
			name := key.(string)
			status := value.(libra.Status)
			code := status.Code.String()
			if _, ok := res[code]; ok {
				res[code] = append(res[code].([]string), name)
			} else {
				res[code] = []string{name}
			}
			a := res[code].([]string)
			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
			res[code] = a
			logrus.Debugf("Failed to execute %v, msg:%v", name, status.Msg)
			return true
		})
		logrus.WithFields(res).Errorf("Job finished. %v/%v succeeded.", reporter.success, reporter.all)
	}
}
