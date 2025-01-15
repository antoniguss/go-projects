package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

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
	RunE:  displayTodos,
}

func displayTodos(cmd *cobra.Command, args []string) error {

	todos, err := todoStorage.GetAllTodos()

	if err != nil {
		return err
	}

	all, _ := cmd.Flags().GetBool("all")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	if all {
		fmt.Fprintln(w, "ID\tDescription\tCompleted\tTime added\tTime completed")
		for _, t := range todos {
			completedAt := "-"
			if t.IsCompleted {
				completedAt = fmt.Sprintf("%v ago", time.Since(t.TimeCompleted.Time).Round(time.Second))
			}
			fmt.Fprintf(w, "%d\t%s\t%v\t%v ago\t%s\n",
				t.ID,
				t.Description,
				t.IsCompleted,
				time.Since(t.TimeAdded).Round(time.Second),
				completedAt,
			)
		}
		return nil
	}

	fmt.Fprintln(w, "ID\tDescription\tTime added")
	for _, t := range todos {
		if !t.IsCompleted {
			fmt.Fprintf(w, "%d\t%s\t%v ago\n", t.ID, t.Description, time.Since(t.TimeAdded).Round(time.Second))
		}
	}

	return nil

}
