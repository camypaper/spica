package judge

import (
	"path/filepath"

	"github.com/Camypaper/libra"

	"github.com/Sirupsen/logrus"
)

/*
Chk checks submission.
*/
func (w SpicaWorker) Chk(target string, submissions []libra.Submission) error {
	checker := *w.Problem.Checker
	lang, err := w.Config.Find(checker)
	if err != nil {
		logrus.WithError(err).Errorf("language for %v does not found.", checker)
		return err
	}
	src := libra.Src{Name: checker, Compile: lang.Compile, Exec: lang.Exec}
	job := libra.ChkJob(src, submissions, target)
	abs, err := filepath.Abs(checker)
	if err != nil {
		logrus.WithError(err).Errorf("checker is not valid.")
		return err
	}
	rel, err := filepath.Rel(".", checker)
	if err != nil {
		logrus.WithError(err).Errorf("checker is not valid.")
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
