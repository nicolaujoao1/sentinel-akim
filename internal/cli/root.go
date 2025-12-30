package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-akim",
	Short: "System health sentinel CLI",
	Long:  "sentinel-akim is a CLI tool to monitor and diagnose system health",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a command. Try --help")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
