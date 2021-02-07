package main

import (
	"fmt"
	"os"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kind/pkg/cluster"
)

var deleteClusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"rm", "del"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Delete a cluster",
	Long:    "The delete sub-commands delete EKS-D clusters.",
	Example: `  # Delete the cluster, the default name is 'ekz'
  ekz delete cluster

  # Delete the 'dev' cluster
  ekz delete cluster dev

  # Delete the 'dev' cluster (alternative syntax)
  ekz delete cluster --name=dev

  # Delete the cluster created by the EKZ provider
  ekz --provider=ekz delete cluster

  # Delete the cluster created by the KIND provider
  ekz --provider=kind delete cluster
`,
	RunE: deleteClusterCmdRun,
}

func init() {
	deleteClusterCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", "kubeconfig", "specify output file to write kubeconfig to")
	deleteClusterCmd.Flags().StringVar(&clusterName, "name", "ekz", "cluster name")

	deleteCmd.AddCommand(deleteClusterCmd)
}

func deleteClusterCmdRun(cmd *cobra.Command, args []string) error {
	logger.Successf("the default provider is: %s", provider)

	// use args[0] as the clusterName
	if len(args) == 1 {
		clusterName = args[0]
	}

	switch provider {
	case "ekz":
		return deleteClusterEKZRun()
	case "kind":
		return deleteClusterKINDRun()
	default:
		return fmt.Errorf("NYI provider: %s", provider)
	}
}

func deleteClusterEKZRun() error {
	containerName := fmt.Sprintf("%s-controller-0", clusterName)
	containerId := script.Var()
	var err error

	err = script.Exec("docker", "ps",
		"-aq",                        // check all containers
		"-f", "name="+containerName). // filter only #{containerName}
		To(containerId)
	if err != nil {
		return errors.Wrapf(err, "failed to run docker ps to check container: %s.", containerName)
	}

	// container does not exist, abort
	if containerId.String() == "" {
		return errors.Errorf("container %s does not exist - cluster deletion aborted", containerName)
	}

	// TODO check if it's the host mode
	networkName, err := getEKZNetworkName(containerName)
	if err != nil {
		return err
	}

	err = script.Exec("docker", "rm",
		"-f", // force delete
		"-v", // remove volume
		containerName).
		Run()
	if err != nil {
		return errors.Wrapf(err, "failed to remove %s", containerName)
	}

	if networkName != "host" {
		// remove bridge after deleting the cluster
		bridgeName := fmt.Sprintf("ekz-%s-bridge", clusterName)
		err = script.Exec("docker", "network", "rm", bridgeName).Run()
		if err != nil {
			return errors.Wrapf(err, "failed to remove bridge: %s", bridgeName)
		}
	}
	return nil
}

func deleteClusterKINDRun() error {
	provider := cluster.NewProvider()
	os.Setenv("KIND_EXPERIMENTAL_DOCKER_NETWORK", fmt.Sprintf("%s-bridge", clusterName))

	return provider.Delete(clusterName, kubeConfigFile)
}
