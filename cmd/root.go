/*
Copyright © 2024 Forest
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan is a todo application.",
	Long:  `Plan will help you get more done in less time. It's designed to be as simple as possible to help you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	rootCmd.PersistentFlags().StringVar(
		&dataFile, "datafile",
		home+string(os.PathSeparator)+".plandos.json", "data file to store todos")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
