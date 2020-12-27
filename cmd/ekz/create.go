package main

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create clusters",
	Long:  "The create sub-commands create EKS-D clusters.",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
