/*
Copyright © 2025 ilhamta27@gmail.com
*/
package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is todo application",
	Long:  `Tri will help you manage your tasks. It's designed to be simple and easy to use.`,
}

var filename string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Init the filename flag
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting home directory: %v\n", err)
	}
	// set the default filename
	filename = path.Join(dirname, "/tri/tri.json")
	rootCmd.PersistentFlags().StringVarP(&filename, "filepath", "f", filename, "todo file (default is $HOME/tri/tri.json)")
}
