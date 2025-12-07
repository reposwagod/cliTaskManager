package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Loading tasks...\n")
		if len(tasks) == 0 {
			fmt.Println("No tasks")
			return
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d tasks:\n", len(tasks))
		}

		for _, task := range tasks {
			status := "In process"
			if task.Completed {
				status = "Completed"
			}
			fmt.Printf("[%d] [%s] %s\n", task.ID, status, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
