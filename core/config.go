package core

import (
	"fmt"
	"path/filepath"
)

/*
Config contains languages and compile timelimits
*/
type Config struct {
	Languages []Lang   `toml:"languages"`
	Timelimit *float64 `toml:"timelimit"`
	Workers   *int     `toml:"workers"`
}

/*
Find a language which extention eqauals to path's one
*/
func (config Config) Find(path string) (Lang, error) {
	ext := filepath.Ext(path)
	for _, v := range config.Languages {
		if v.Ext == ext {
			return v, nil
		}
	}
	return Lang{}, fmt.Errorf("language for %s was not found", ext)
}

func (config Config) String() string {
	return fmt.Sprintf("[Languages:%v, Timelimit:%.1fsec, Workers:%v]", config.Languages, *config.Timelimit, *config.Workers)
}
