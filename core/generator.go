package core

import (
	"fmt"

	"github.com/camypaper/libra"

	"github.com/sirupsen/logrus"
)

/*
Generator setting.
*/
type Generator struct {
	Name string `toml:"name"`
	Cnt  int    `toml:"cnt"`
}

/*
ToJob converts to GenJob
*/
func (generator Generator) ToJob(config Config) (libra.Job, error) {
	lang, err := config.Find(generator.Name)
	if err != nil {
		logrus.WithError(err).Errorf("language for %v does not found.", generator.Name)
		return nil, err
	}
	src := libra.Src{Name: generator.Name, Compile: lang.Compile, Exec: lang.Exec}
	return libra.GenJob(src, generator.Cnt), nil
}

func (generator Generator) String() string {
	return fmt.Sprintf("[Name:%v, Cnt:%v]", generator.Name, generator.Cnt)
}
