package cmd

import (
	"github.com/Camypaper/spica/core"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	ret := &cobra.Command{
		Use:   "add",
		Short: "Add fact to the problem",
		Long:  "Add fact to the problem",
	}
	ret.AddCommand(addGenCmd(), addAnsCmd(), setValCmd(), setChkCmd(), setSolCmd())
	return ret
}

func addGenCmd() *cobra.Command {
	var cnt int
	ret := &cobra.Command{
		Use:   "gen",
		Short: "Add generator to the problem",
		Long:  "Add generator to the problem",
		Run: func(cmd *cobra.Command, args []string) {
			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			for _, v := range args {
				problem.Generators = append(problem.Generators, core.Generator{Name: v, Cnt: cnt})
			}
			err = core.SaveProblem(problem)
			if err != nil {
				logrus.Fatal("Failed to save problem")
			}
		},
		Args: cobra.MinimumNArgs(1),
	}
	ret.Flags().IntVarP(&cnt, "cnt", "c", 0, "size of generated files")
	return ret
}

func addAnsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ans",
		Short: "Add answer to the problem",
		Long:  "Add answer to the problem",
		Run: func(cmd *cobra.Command, args []string) {
			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			for _, v := range args {
				problem.Answers = append(problem.Answers, v)
			}
			err = core.SaveProblem(problem)
			if err != nil {
				logrus.Fatal("Failed to save problem")
			}
		},
		Args: cobra.MinimumNArgs(1),
	}
}

func setValCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "val",
		Short: "Set validator to the problem",
		Long:  "Set validator to the problem",
		Run: func(cmd *cobra.Command, args []string) {
			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			problem.Validator = &args[0]
			err = core.SaveProblem(problem)
			if err != nil {
				logrus.Fatal("Failed to save problem")
			}
		},
		Args: cobra.ExactArgs(1),
	}
}

func setSolCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sol",
		Short: "Set solution to the problem",
		Long:  "Set solution to the problem",
		Run: func(cmd *cobra.Command, args []string) {
			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			problem.Solution = &args[0]
			err = core.SaveProblem(problem)
			if err != nil {
				logrus.Fatal("Failed to save problem")
			}
		},
		Args: cobra.ExactArgs(1),
	}
}

func setChkCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "chk",
		Short: "Set Checker to the problem",
		Long:  "Set Checker to the problem",
		Run: func(cmd *cobra.Command, args []string) {
			problem, err := core.LoadProblem()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to load problem.")
			}
			logrus.WithField("problem", problem).Info("Succeeded to load problem.")
			problem.Checker = &args[0]
			err = core.SaveProblem(problem)
			if err != nil {
				logrus.Fatal("Failed to save problem")
			}
		},
		Args: cobra.ExactArgs(1),
	}
}
