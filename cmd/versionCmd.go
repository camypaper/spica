package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of spica",
		Long:  `All software has versions. This is spica's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("spica: competetive programming setting tools v0.0.1")
		},
	}
}
