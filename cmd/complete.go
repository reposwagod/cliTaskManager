package cmd

import (
	"fmt"
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
			return
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
				saveTasks()
				fmt.Printf("Completed task: %s\n", tasks[i].Title)
				return
			}
		}
		fmt.Println("Task not found")
	},
}
