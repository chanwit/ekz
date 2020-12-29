package main

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List clusters",
	Long:    "The list sub-commands list EKS-D clusters. Currently it acts as a shortcut for the 'list clusters' command.",
	RunE:    listClusterCmdRun, // also default to the "list cluster"
}

func init() {
	rootCmd.AddCommand(listCmd)
}
