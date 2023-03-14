package cmd

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI task manager.",
}
