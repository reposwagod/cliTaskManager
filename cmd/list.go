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
			fmt.Printf("ID: [%d]\nStatus: %s\nName: %s\nPriority: %s\n\n", task.ID, status, task.Title, task.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
