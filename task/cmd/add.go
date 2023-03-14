package cmd

import (
	"fmt"

	"strings"

	"github.com/EdwinVesga/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added task: %q in the task list.\n", task)
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		}
	},
}

func init() {
	RootCommand.AddCommand(addCmd)
}
