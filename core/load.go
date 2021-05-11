package core

import (
	"github.com/sirupsen/logrus"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

/*
LoadProblem loads problem.
*/
func LoadProblem() (Problem, error) {
	problem := Problem{}
	v := viper.New()
	v.SetConfigName("problem")
	v.AddConfigPath(".")

	v.SetDefault("timelimit", 2.0)
	err := v.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Error("Problem config not found.")
		return problem, err
	}
	err = v.Unmarshal(&problem)
	if err != nil {
		logrus.WithError(err).Error("Problem config is not valid")
	}

	return problem, nil
}

/*
LoadConfig loads config.
*/
func LoadConfig() (Config, error) {
	config := Config{}
	tl := 30.0
	workers := 4
	langs := []Lang{}
	langs = append(langs, Lang{Ext: ".cpp", Compile: "g++ $SRC", Exec: "./a.exe"})
	langs = append(langs, Lang{Ext: ".cc", Compile: "g++ $SRC", Exec: "./a.exe"})
	langs = append(langs, Lang{Ext: ".cs", Compile: "mcs -r:System.Numerics.dll $SRC", Exec: "mono ./a.exe"})
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if home, err := homedir.Dir(); err == nil {
		v.AddConfigPath(home)
	}
	v.SetDefault("timelimit", tl)
	v.SetDefault("workers", workers)
	err := v.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Warn("Config file not found. spica will use default config.")
		return Config{Timelimit: &tl, Workers: &workers, Languages: langs}, nil
	}
	err = v.Unmarshal(&config)
	if err != nil {
		logrus.WithError(err).Error("Config file is not valid")
	}

	return config, nil
}
