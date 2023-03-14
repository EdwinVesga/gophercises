package cmd

import (
	"fmt"
	"strconv"

	"github.com/EdwinVesga/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do the tasks with the ids passed.",
	Run: func(cmd *cobra.Command, args []string) {
		var doList []int
		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid do argument.")
			} else {
				doList = append(doList, i)
			}
		}
		fmt.Printf("Do items: %+v\n", doList)
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}
		for _, id := range doList {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}

		}
	},
}

func init() {
	RootCommand.AddCommand(doCmd)
}
