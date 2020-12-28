package main

import (
	"fmt"
	"time"

	"github.com/chanwit/script"
	"github.com/spf13/cobra"
)

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a cluster",
	Long:  "The create sub-commands create EKS-D clusters.",
	RunE:  createClusterCmdRun,
}

var (
	eksdVersion    string
	kubeConfigFile string
)

func init() {
	createClusterCmd.Flags().StringVar(&eksdVersion, "eksd-version", "v1.18.9-eks-1-18-1", "specify a version of EKS-D")
	createClusterCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", "kubeconfig", "specify output file to write kubeconfig to")

	createCmd.AddCommand(createClusterCmd)
}

func createClusterCmdRun(cmd *cobra.Command, args []string) error {
	switch provider {
	case "ekz":
		return createClusterEKZ()
	case "kind":
		return createClusterKIND()
	default:
		return fmt.Errorf("NYI provider: %s", provider)
	}
}

func waitForNodeStarted(nodeName string) {
	for {
		name := script.Var()
		script.Exec("kubectl", "--kubeconfig="+kubeConfigFile, "get", "nodes", "-ojsonpath={.items[0].metadata.name}").To(name)
		if name.String() == nodeName {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func waitForNodeReady() {
	for {
		status := script.Var()
		script.Exec("kubectl", "--kubeconfig="+kubeConfigFile, "get", "nodes", "-ojsonpath={.items[0].status.conditions[-1].type}={.items[0].status.conditions[-1].status}").To(status)
		if status.String() == "Ready=True" {
			break
		}
		time.Sleep(2 * time.Second)
	}
}
