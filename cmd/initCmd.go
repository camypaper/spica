package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the problem",
		Long:  "Initialize the problem with problem.toml",
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.New()
			v.SetConfigName("problem")
			v.AddConfigPath(".")
			err := v.ReadInConfig()
			if err == nil {
				logrus.Fatal("Problem config already exists.")
				return
			}

			err = v.WriteConfigAs("problem.toml")
			if err != nil {
				logrus.WithError(err).Fatalf("Failed to create config")
				return
			}
			logrus.Info("Succeeded to create config.")
		},
	}
}
