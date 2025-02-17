/*
Copyright © 2025 ilhamta27@gmail.com
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

var (
	doneOpt bool
	allOpt  bool
)

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(filename)
	if err != nil {
		log.Printf("Error reading todo items: %v\n", err)
		return
	}
	printItems(items)
}

func printItems(items []todo.Item) {
	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	fmt.Fprintln(w, "No\tPriority\tTodo\tStatus")
	fmt.Fprintln(w, "__\t_______\t_______\t__________")
	for _, item := range items {
		if allOpt || doneOpt == item.Done {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.PrettyStatus())
		}
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show only done items")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all items (including done)")
}
