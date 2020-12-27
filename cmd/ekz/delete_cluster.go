package main

import (
	"fmt"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var deleteClusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"rm", "del"},
	Short:   "Delete a cluster",
	Long:    "The create sub-commands create EKS-D clusters.",
	RunE:    deleteClusterCmdRun,
}

func init() {
	deleteClusterCmd.Flags().StringVar(&provider, "provider", "ekz", "cluster provider (ekz, kind)")

	deleteCmd.AddCommand(deleteClusterCmd)
}

func deleteClusterCmdRun(cmd *cobra.Command, args []string) error {
	switch provider {
	case "ekz":
		return deleteClusterEKZRun(cmd, args)
	case "kind":
		return deleteClusterKINDRun(cmd, args)
	default:
		return fmt.Errorf("NYI provider: %s", provider)
	}
}

func deleteClusterEKZRun(cmd *cobra.Command, args []string) error {
	containerName := "ekz-controller-0"
	containerId := script.Var()
	var err error

	err = script.Exec("docker", "ps", "-aq", "-f", "name="+containerName).To(containerId)
	if err != nil {
		return errors.Wrapf(err, "failed to run docker ps to check container: %s.", containerName)
	}

	// container does not exist, abort
	if containerId.String() == "" {
		return errors.Errorf("container %s does not exist. cluster deletion aborted.", containerName)
	}

	err = script.Exec("docker", "rm", "-f", containerName).Run()
	if err != nil {
		return errors.Wrapf(err, "failed to remove %s", containerName)
	}

	return nil
}

func deleteClusterKINDRun(cmd *cobra.Command, args []string) error {
	return nil
}
