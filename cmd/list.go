/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ilhamtubagus/tri/todo"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List will display all the todos in your todo list.`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(filename)
	if err != nil {
		log.Printf("Error reading todo items: %v\n", err)
		return
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, item := range items {
		fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
