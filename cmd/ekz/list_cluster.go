package main

import (
	"github.com/spf13/cobra"
)

var listClusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"clusters"},
	Short:   "List cluster",
	Long:    "",
	RunE:    listClusterCmdRun,
}

func init() {
	listCmd.AddCommand(listClusterCmd)
}

// How to heuristically detect a kind / ekz cluster
// "io.x-k8s.kind.cluster": "ekz"

// How to heuristically detect a ekz cluster
// io.x-k8s.ekz.cluster=ekz

func listClusterCmdRun(cmd *cobra.Command, args []string) error {
	return nil
}
