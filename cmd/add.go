package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var priority string

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		if verbose {
			fmt.Fprintf(os.Stderr, "Adding task: %s\n", description)
			defer fmt.Fprintf(os.Stderr, "Task ID: %d\n", len(tasks))
		}

		task := Task{
			ID:        len(tasks) + 1,
			Title:     args[0],
			Completed: false,
			Priority:  priority,
		}
		tasks = append(tasks, task)
		saveTasks()
		fmt.Printf("Added task: %s\n", args[0])
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "medium", "Add priority to your task")
	rootCmd.AddCommand(addCmd)
}
