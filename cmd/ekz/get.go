package main

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get properties of an EKS-D cluster",
	Long:  "Get sub-commands get properties of an EKS-D cluster.",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
