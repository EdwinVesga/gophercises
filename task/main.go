package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EdwinVesga/gophercises/task/cmd"
	"github.com/EdwinVesga/gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	must(cmd.RootCommand.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
