package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/camypaper/libra"
	"github.com/camypaper/spica/core"
	"github.com/camypaper/spica/judge"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func publishCmd() *cobra.Command {
	var isStrict bool
	var target string
	var all bool
	ret := &cobra.Command{
		Use:   "publish",
		Short: "Publish the problem",
		Long:  "Publish the problem",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := core.LoadConfig()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load config.")
			}
			logrus.WithField("config", config).Info("Succeeded to load config.")

			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			worker := judge.SpicaWorker{Config: config, Problem: problem}
			worker.Worker = libra.ConcurrentWorker{Num: *config.Workers}

			tests, err := generate(worker)
			if isStrict && err != nil {
				logrus.WithError(err).Fatalf("Failed to generate.")
			}
			tests, err = validate(worker, tests)
			if isStrict && err != nil {
				logrus.WithError(err).Fatalf("Failed to validate.")
			}
			tests, err = execute(worker, tests)
			if isStrict && err != nil {
				logrus.WithError(err).Fatalf("Failed to execute.")
			}
			if err := publishTestCase(target, tests); err != nil {
				logrus.WithError(err).Fatalf("Failed to publish.")
			}
			if all {
				for _, v := range problem.Answers {
					submissions := check(worker, v, tests)
					publishSubmission(filepath.Join(target, "ans", v), submissions)
				}
			}
		},
	}
	ret.Flags().BoolVarP(&isStrict, "strict", "s", true, "strict publish")
	ret.Flags().BoolVarP(&all, "all", "a", false, "publish with all submissions")
	ret.Flags().StringVarP(&target, "to", "t", ".", "target directory")

	return ret
}

func publishTestCase(to string, tests []core.Testcase) error {
	in := filepath.Join(to, "in")
	if err := os.RemoveAll(in); err != nil {
		return err
	}
	out := filepath.Join(to, "out")
	if err := os.RemoveAll(out); err != nil {
		return err
	}

	for _, v := range tests {
		if v.In == nil || v.Out == nil {
			return fmt.Errorf("%v: input or output is nil", v.Name)
		}
		name := v.Name + ".txt"
		inFile, err := create(filepath.Join(in, name))
		if err != nil {
			return err
		}

		outFile, err := create(filepath.Join(out, name))
		if err != nil {
			return err
		}
		if _, err := inFile.WriteString(*v.In); err != nil {
			return err
		}
		if _, err := outFile.WriteString(*v.Out); err != nil {
			return err
		}
	}
	return nil
}

func publishSubmission(to string, tests []libra.Submission) error {
	ans := filepath.Join(to)
	if err := os.RemoveAll(ans); err != nil {
		return err
	}
	for _, v := range tests {
		name := v.Name + ".txt"
		ansFile, err := create(filepath.Join(ans, name))
		if err != nil {
			return err
		}
		if _, err := ansFile.WriteString(v.Status.Msg); err != nil {
			return err
		}
	}
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
