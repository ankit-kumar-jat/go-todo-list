package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks in your todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ok, from tasks lists")
		return nil
	},
}
