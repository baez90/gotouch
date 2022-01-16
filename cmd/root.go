/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/denizgursoy/gotouch/prompt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotouch",
	Short: "Helps you create new Go Projects",
	Long:  `Tag`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing the command")
		createNewProject()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotouch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createNewProject() {

	httpServerOptions := make([]string, 0)
	httpServerOptions = append(httpServerOptions, "echo")
	httpServerOptions = append(httpServerOptions, "mux")
	httpServerOptions = append(httpServerOptions, "fiber")

	selection := prompt.AskForSelection(prompt.Definition{
		ErrorText: "Please select a HTTP Framework",
		Direction: "Select one from the available options",
	}, httpServerOptions)

	fmt.Println("user selected", selection)
}
