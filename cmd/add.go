/*
Copyright © 2025 ilhamta27@gmail.com
*/
package cmd

import (
	"github.com/ilhamtubagus/tri/todo"
	"github.com/spf13/cobra"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo with the given description into todo list.`,
	Run:   addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(filename)
	if err != nil {
		log.Printf("Error reading todo items: %v\n", err)
	}

	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(
			items,
			item,
		)
	}

	err = todo.SaveItems(filename, items)
	if err != nil {
		log.Printf("Error saving todo items: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority (1: high, 3: low, 2: default/normal)")
}
