/*
Copyright © 2024 Forest
*/
package cmd

import (
	"log"
	"plan/todo"

	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use add to add new todo tasks.",
	Long: "Add is a command that allows you to add new todo task. It provides multiple flags. Run `plan help add` to find out more.",
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
