package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().BoolP("undo", "u", false, "Undo the completion of a TODO")

}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a TODO as completed",
	Long:  `Mark a TODO with the given ID as completed, or undo if --undo is specified`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := strconv.Atoi(args[0])

		if err != nil {
			return fmt.Errorf("invalid ID: %s", args[0])
		}

		undo, _ := cmd.Flags().GetBool("undo")

		if err := todoStorage.SetCompleted(id, !undo); err != nil {
			return fmt.Errorf("invalid ID: %s", err)
		}

		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			if undo {
				fmt.Printf("Undoing completion of TODO %d\n", id)
			} else {
				fmt.Printf("Marking TODO %d as completed\n", id)
			}
		}

		return nil
	},
}
