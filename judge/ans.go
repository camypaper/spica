package judge

import (
	"path/filepath"

	"github.com/Camypaper/libra"
	"github.com/Camypaper/spica/core"
	"github.com/Sirupsen/logrus"
)

/*
Ans generates output.
*/
func (w SpicaWorker) Ans(answer string, testcases []core.Testcase) error {
	lang, err := w.Config.Find(answer)
	if err != nil {
		logrus.WithError(err).Errorf("language for %v does not found.", answer)
		return err
	}

	src := libra.Src{Name: answer, Compile: lang.Compile, Exec: lang.Exec}
	stdins := []libra.Input{}
	for _, v := range testcases {
		input, err := v.ToInput()
		if err != nil {
			logrus.WithField("target", v.Name).WithError(err).Errorf("Testcase Not Found.")
		} else {
			stdins = append(stdins, input)
		}
	}

	job := libra.AnsJob(src, stdins)
	abs, err := filepath.Abs(answer)
	if err != nil {
		logrus.WithError(err).Errorf("answer is not valid.")
		return err
	}
	rel, err := filepath.Rel(".", answer)
	if err != nil {
		logrus.WithError(err).Errorf("answer is not valid.")
		return err
	}
	f := func() { w.ExecTask(job, w.Context) }
	g := func() error {
		withSrc(abs, rel, f)
		return nil
	}
	err = withTemp(w.Problem.Resources, g)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to execute %v.", filepath.Base(src.Name))
	}
	return err
}
