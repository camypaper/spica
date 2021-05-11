package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	var verbose bool
	cmd := &cobra.Command{
		Use:   "spica",
		Short: "Testing tool for competetive programming writer",
		Long:  `Testing tool for competetive programming writer`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
			logrus.SetOutput(colorable.NewColorableStdout())
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}
		},
	}
	cmd.AddCommand(versionCmd(), initCmd(), testCmd(), publishCmd(), addCmd())
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	return cmd
}

/*
Execute spica.
*/
func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
