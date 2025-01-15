package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Long:  `Add a new TODO with the given description`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args[0]) == 0 {
			return fmt.Errorf("TODO description cannot be empty")
		}
		id := todoStorage.AddTodo(args[0])

		//Check if debug flag is set
		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			fmt.Printf("Added TODO with id %d\n", id)
		}
		return nil

	},
}
