package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete id",
	Short: "Removes a task for the todo list by it's id",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ok, from tasks delete")
		return nil
	},
}
