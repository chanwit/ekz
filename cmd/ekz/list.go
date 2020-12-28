package main

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List clusters",
	Long:    "The list sub-commands list EKS-D clusters.",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
