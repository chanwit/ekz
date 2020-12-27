package main

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete clusters",
	Long:  "The delete sub-commands delete EKS-D clusters.",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
