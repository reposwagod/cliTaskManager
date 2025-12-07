package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) == 0 {
			fmt.Println("No tasks")
			return
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
