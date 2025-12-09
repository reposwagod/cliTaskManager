package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const ver = "1.0.4"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Current version: %s", ver)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
