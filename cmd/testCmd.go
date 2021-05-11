package cmd

import (
	"github.com/camypaper/libra"
	"github.com/camypaper/spica/core"
	"github.com/camypaper/spica/judge"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func testCmd() *cobra.Command {
	var isConcurrent bool
	ret := &cobra.Command{
		Use:   "test",
		Short: "Test the problem",
		Long:  "Test the problem",
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
			if isConcurrent {
				worker.Worker = libra.ConcurrentWorker{Num: *config.Workers}
			} else {
				worker.Worker = libra.SequentialWorker{}
			}
			tests, err := generate(worker)
			tests, err = validate(worker, tests)
			tests, err = execute(worker, tests)
			for _, v := range problem.Answers {
				check(worker, v, tests)
			}
		},
	}
	ret.Flags().BoolVarP(&isConcurrent, "concurrent", "c", false, "concurrent execution")

	return ret
}
