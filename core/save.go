package core

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

/*
SaveProblem saves problem.toml to working directory
*/
func SaveProblem(problem Problem) error {
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
	err = tosave.WriteConfigAs("problem.toml")
	if err != nil {
		logrus.WithError(err).Error("Failed to write problem")
		return err
	}
	err = generate(problem)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate files")
		return err
	}
	return nil
}

func generate(problem Problem) error {
	target := []string{}
	target = append(target, problem.Answers...)
	if problem.Checker != nil {
		target = append(target, *problem.Checker)
	}
	for _, v := range problem.Generators {
		target = append(target, v.Name)
	}
	if problem.Solution != nil {
		target = append(target, *problem.Solution)
	}
	if problem.Validator != nil {
		target = append(target, *problem.Validator)
	}
	for _, v := range target {
		if v != "" {
			_, err := create(v)
			if err != nil {
				return err
			}
		}
	}
	logrus.WithField("problem", problem).Info("Succeeded to save problem")
	return nil
}

func create(path string) (*os.File, error) {
	base := filepath.Dir(path)
	err := os.MkdirAll(base, 0755)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
}
