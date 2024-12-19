package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   `add "my new task item"`,
	Short: "Add a new task to the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		fmt.Println(args[0])
		return nil
	},
}
