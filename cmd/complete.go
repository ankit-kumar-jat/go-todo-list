package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete id",
	Short: "Set a task as being completed",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ok, from tasks complete")
		return nil
	},
}
