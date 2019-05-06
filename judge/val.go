package judge

import (
	"path/filepath"

	"github.com/Camypaper/libra"
	"github.com/Camypaper/spica/core"

	"github.com/Sirupsen/logrus"
)

/*
Val validates input.
*/
func (w SpicaWorker) Val(testcases []core.Testcase) error {
	validator := *w.Problem.Validator
	lang, err := w.Config.Find(validator)
	if err != nil {
		logrus.WithError(err).Errorf("language for %v does not found.", validator)
		return err
	}

	src := libra.Src{Name: validator, Compile: lang.Compile, Exec: lang.Exec}
	stdins := []libra.Input{}
	for _, v := range testcases {
		input, err := v.ToInput()
		if err != nil {
			logrus.WithField("target", v.Name).WithError(err).Errorf("Testcase Not Found.")
		} else {
			stdins = append(stdins, input)
		}
	}

	job := libra.ValJob(src, stdins)
	abs, err := filepath.Abs(validator)
	if err != nil {
		logrus.WithError(err).Errorf("Validator is not valid.")
		return err
	}
	rel, err := filepath.Rel(".", validator)
	if err != nil {
		logrus.WithError(err).Errorf("Validator is not valid.")
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
