/*
Copyright © 2025 ilhamta27@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/ilhamtubagus/tri/todo"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark todo item as done",
	Aliases: []string{"do"},
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	i, err := strconv.Atoi(args[0])
	if err != nil || i < 1 {
		log.Fatalln(args[0], "is not a valid label", err)
	}
	items, err := todo.ReadItems(filename)
	if err != nil {
		log.Fatalln("Error reading todo items: ", err)
		return
	}

	if i > len(items) {
		log.Fatalln("Invalid label. Label should be between 1 and", len(items))
	}

	items[i-1].Done = true
	fmt.Println("Marked todo item as done:", items[i-1].Text)

	sort.Sort(todo.ByPri(items))
	err = todo.SaveItems(filename, items)
	if err != nil {
		log.Fatalln("Error saving todo items: ", err)
	}
	printItems(items)
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
