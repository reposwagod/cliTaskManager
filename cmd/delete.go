package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			if verbose {
				fmt.Println("Try ./task list or task list for more information")
			}
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				if verbose {
					fmt.Fprintf(os.Stderr, "Deleting task: %d\n", id)
					defer fmt.Fprintf(os.Stderr, "Remaining tasks: %d\n", len(tasks))
				}
				tasks = append(tasks[:i], tasks[i+1:]...)
				saveTasks()
				fmt.Printf("Deleted task: %s\n", task.Title)
				return
			}
		}
		fmt.Println("Task not found")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
