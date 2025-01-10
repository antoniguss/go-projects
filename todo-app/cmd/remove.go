package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a TODO",
	Long:    `Remove a TODO with the given ID`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := strconv.Atoi(args[0])

		if err != nil {
			return fmt.Errorf("invalid ID: %s", args[0])
		}

		if err := todoManager.RemoveTodo(id); err != nil {
			return fmt.Errorf("invalid ID: %s", err)
		}

		return nil
	},
}
