package judge

import (
	"path/filepath"

	"github.com/camypaper/spica/core"
	"github.com/sirupsen/logrus"
)

/*
Gen generates input.
*/
func (w SpicaWorker) Gen(generator core.Generator) error {
	job, err := generator.ToJob(w.Config)
	if err != nil {
		logrus.WithError(err).Errorf("Generator is not valid.")
		return err
	}
	abs, err := filepath.Abs(generator.Name)
	if err != nil {
		logrus.WithError(err).Errorf("Generator is not valid.")
		return err
	}
	rel, err := filepath.Rel(".", generator.Name)
	if err != nil {
		logrus.WithError(err).Errorf("Generator is not valid.")
		return err
	}
	f := func() { w.ExecTask(job, w.Context) }
	g := func() error {
		withSrc(abs, rel, f)
		return nil
	}
	err = withTemp(w.Problem.Resources, g)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to execute %v.", filepath.Base(generator.Name))
	}
	return err
}
