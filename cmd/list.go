package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var sortby string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Loading tasks...\n\n")
		if len(tasks) == 0 {
			fmt.Println("No tasks")
			return
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d tasks:\n", len(tasks))
		}

		if sortby != "" {
			sortTasks(tasks, sortby)
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

func sortTasks(tasks []Task, sortBy string) {
	switch strings.ToLower(sortBy) {
	case "name", "title":
		sort.Slice(tasks, func(i, j int) bool {
			return strings.ToLower(tasks[i].Title) < strings.ToLower(tasks[j].Title)
		})

	case "priority":
		priorityOrder := map[string]int{
			"high":   1,
			"medium": 2,
			"low":    3,
		}

		sort.Slice(tasks, func(i, j int) bool {
			prioI := priorityOrder[strings.ToLower(tasks[i].Priority)]
			prioJ := priorityOrder[strings.ToLower(tasks[j].Priority)]
			return prioI < prioJ
		})

	case "status", "completed":
		sort.Slice(tasks, func(i, j int) bool {
			if tasks[i].Completed && !tasks[j].Completed {
				return false
			}

			if !tasks[i].Completed && tasks[j].Completed {
				return true
			}

			return tasks[i].ID < tasks[j].ID
		})
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&sortby, "sortby", "s", "ID", "Sort list by (ID, Name, Priority)")
	rootCmd.AddCommand(listCmd)
}
