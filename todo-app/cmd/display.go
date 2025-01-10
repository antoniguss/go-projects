package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(displayCmd)
	displayCmd.Flags().BoolP("all", "a", false, "Show completed TODOs")
}

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display TODOs",
	Long:  `Display a list of TODOs, optionally showing completed items`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		todoManager.DisplayTodos(all)
	},
}
