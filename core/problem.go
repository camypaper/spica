package core

import (
	"bytes"
	"fmt"

	"github.com/Sirupsen/logrus"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

/*
Problem setting
*/
type Problem struct {
	Timelimit  *float64    `toml:"timelimit"`
	Generators []Generator `toml:"generators"`
	Answers    []string    `toml:"answers"`
	Validator  *string     `toml:"validator"`
	Checker    *string     `toml:"checker"`
	Solution   *string     `toml:"solution"`
	Resources  []string    `toml:"resources"`
}

func (problem Problem) save(path string) error {
	data, err := toml.Marshal(problem)
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal problem")
		return err
	}
	tosave := viper.New()
	tosave.SetConfigType("toml")
	err = tosave.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logrus.WithError(err).Error("Failed to read problem")
		return err
	}
	err = tosave.WriteConfigAs(path)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write problem")
	}
	return err
}

func (problem Problem) String() string {
	f := func(s *string) string {
		if s == nil {
			return "<nil>"
		}
		return *s
	}

	return fmt.Sprintf("[Timelimit:%.1fsec, Generators:%v, Answers:%v, Validator:%+v, Checker:%+v, Solution: %+v, Resources: %v]",
		*problem.Timelimit, problem.Generators, problem.Answers, f(problem.Validator), f(problem.Checker), f(problem.Solution), problem.Resources)
}
