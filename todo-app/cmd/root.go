package cmd

import (
	"log"
	"os"

	"github.com/antoniguss/go-projects/todo-app/storage"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-app",
	Short: "A simple todo app",
	Long:  "A simple CLI todo app I'm building in GoLang.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var todoStorage storage.Storage

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	err = todoStorage.Close()
	if err != nil {
		log.Print(err)
	}

}

func init() {
	var err error
	todoStorage, err = storage.NewStorage("sqlite", "./todos.db")

	if err != nil {
		log.Fatal(err)
	}

	err = todoStorage.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "log out all the debug information")
}
