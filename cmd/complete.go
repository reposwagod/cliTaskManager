package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark task as complete",
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

		for i := range tasks {
			if tasks[i].ID == id {
				if verbose {
					fmt.Fprintf(os.Stderr, "Completing task ID: %d\n", id)
				}
				tasks[i].Completed = true
				saveTasks()
				fmt.Printf("Completed task: %s\n", tasks[i].Title)
				return
			}
		}
		fmt.Println("Task not found")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
