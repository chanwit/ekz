package main

import (
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load artifacts into the cluster",
	Long:  "The load sub-commands load artifacts into the EKS-D clusters.",
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
