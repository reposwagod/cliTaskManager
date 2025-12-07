package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	verbose  bool
	home, _  = os.UserHomeDir()
	tasks    []Task
	taskFile = filepath.Join(home, "cliTaskManager", "tasks.json")
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Cli Task Manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadTasks()
	},
}

func loadTasks() {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		println(err)
		tasks = []Task{}
		return
	}
	json.Unmarshal(data, &tasks)
}

func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		println(err)
		return
	}
	os.WriteFile(taskFile, data, 0644)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed output")
}
