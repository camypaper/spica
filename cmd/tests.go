package cmd

import (
	"fmt"

	"path/filepath"
	"sort"

	"github.com/Camypaper/libra"
	"github.com/Camypaper/spica/core"
	"github.com/Camypaper/spica/judge"
	"github.com/Sirupsen/logrus"
)

func generate(worker judge.SpicaWorker) ([]core.Testcase, error) {
	logrus.Info("Generate inputs start.")
	out := &core.OutputReporter{}
	worker.Context = libra.WorkerContext{
		Reporter:         core.NewMultiReporter(&core.StdoutReporter{}, out),
		InitializeRunner: libra.Runner{TL: *worker.Config.Timelimit},
		Runner:           libra.Runner{TL: *worker.Problem.Timelimit},
	}
	all := 0
	for _, v := range worker.Problem.Generators {
		all += v.Cnt
		worker.Gen(v)
	}
	ret := []core.Testcase{}
	for k, v := range out.Get() {
		s := v.Msg
		ret = append(ret, core.Testcase{Name: k, In: &s})
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Name < ret[j].Name })
	if all != len(ret) {
		return ret, fmt.Errorf("%v/%v succeeded", len(ret), all)
	}
	return ret, nil
}

func validate(worker judge.SpicaWorker, testcases []core.Testcase) ([]core.Testcase, error) {
	if worker.Problem.Validator == nil {
		logrus.Infof("Validation Skipped.")
		return testcases, nil
	}
	logrus.Info("Validation start")
	out := &core.OutputReporter{}
	worker.Context = libra.WorkerContext{
		Reporter:         core.NewMultiReporter(&core.StdoutReporter{}, out),
		InitializeRunner: libra.Runner{TL: *worker.Config.Timelimit},
		Runner:           libra.Runner{TL: *worker.Problem.Timelimit},
	}
	worker.Val(testcases)

	find := func(a []core.Testcase, key string) *core.Testcase {
		for _, v := range a {
			if v.Name == key {
				return &v
			}
		}
		return nil
	}
	ret := []core.Testcase{}
	for k := range out.Get() {
		res := find(testcases, k)
		if res != nil {
			ret = append(ret, *res)
		}
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Name < ret[j].Name })
	if len(testcases) != len(ret) {
		return ret, fmt.Errorf("%v/%v succeeded", len(ret), len(testcases))
	}
	return ret, nil
}

func execute(worker judge.SpicaWorker, testcases []core.Testcase) ([]core.Testcase, error) {
	if worker.Problem.Solution == nil {
		logrus.Infof("Solution not found.")
		return testcases, nil
	}
	logrus.Info("Refrun start.")
	out := &core.OutputReporter{}
	worker.Context = libra.WorkerContext{
		Reporter:         core.NewMultiReporter(&core.StdoutReporter{}, out),
		InitializeRunner: libra.Runner{TL: *worker.Config.Timelimit},
		Runner:           libra.Runner{TL: *worker.Problem.Timelimit},
	}
	worker.Ans(*worker.Problem.Solution, testcases)

	ret := []core.Testcase{}
	for k, v := range out.Get() {
		res := find(testcases, k)
		if res != nil {
			s := v.Msg
			res.Out = &s
			ret = append(ret, *res)
		}
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Name < ret[j].Name })
	if len(testcases) != len(ret) {
		return ret, fmt.Errorf("%v/%v succeeded", len(ret), len(testcases))
	}
	return ret, nil
}

func check(worker judge.SpicaWorker, target string, testcases []core.Testcase) []libra.Submission {
	logrus.Infof("Test to %v", filepath.Base(target))
	out := &core.OutputReporter{All: true}
	worker.Context = libra.WorkerContext{
		Reporter:         out,
		InitializeRunner: libra.Runner{TL: *worker.Config.Timelimit},
		Runner:           libra.Runner{TL: *worker.Problem.Timelimit},
	}
	worker.Ans(target, testcases)
	submissions := []libra.Submission{}
	for k, v := range out.Get() {
		res := find(testcases, k)
		if res != nil {
			submissions = append(submissions, libra.Submission{Name: k, In: *res.In, Ans: *res.Out, Status: v})
		}
	}
	if worker.Problem.Checker == nil {
		logrus.Infof("Checker not found.")
		return submissions
	}
	//worker reset
	worker.Context = libra.WorkerContext{
		Reporter:         core.NewMultiReporter(&core.StdoutReporter{}),
		InitializeRunner: libra.Runner{TL: *worker.Config.Timelimit},
		Runner:           libra.Runner{TL: *worker.Problem.Timelimit},
	}
	worker.Chk(target, submissions)
	return submissions
}

func find(a []core.Testcase, key string) *core.Testcase {
	for _, v := range a {
		if v.Name == key {
			return &v
		}
	}
	return nil
}
