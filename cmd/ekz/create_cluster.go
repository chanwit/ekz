package main

import (
	"fmt"
	"time"

	"github.com/chanwit/script"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a cluster",
	Long:  "The create sub-commands create EKS-D clusters.",
	Example: `  # Create an EKS-D cluster with the default provider
  # The KubeConfig file will be written to $PWD/kubeconfig.
  ekz create cluster

  # Create an EKS-D cluster with the EKZ provider
  # This command creates an EKS-D-compatible K0s-based cluster.
  ekz --provider=ekz create cluster

  # Create an EKS-D cluster with the KIND provider
  # This command creates an EKS-D-compatible KIND cluster.
  ekz --provider=kind create cluster

  # Create an EKS-D cluster and write KubeConfig to $HOME/.kube/config
  ekz create cluster -o $HOME/.kube/config

  # Create EKS-D cluster with a specific version of EKS-D
  ekz create --eksd-version=v1.18.9-eks-1-18-1 cluster 
`,
	RunE: createClusterCmdRun,
}

var (
	eksdVersion    string
	kubeConfigFile string
	clusterName    string
)

func init() {
	createClusterCmd.Flags().StringVar(&eksdVersion, "eksd-version", "v1.18.9-eks-1-18-1", "specify a version of EKS-D")
	createClusterCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", clientcmd.RecommendedHomeFile, "specify output file to write kubeconfig to")
	createClusterCmd.Flags().StringVar(&clusterName, "name", "ekz", "cluster name")

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
		script.Exec("kubectl", "--kubeconfig="+kubeConfigFile,
			"get", "nodes",
			"-ojsonpath={.items[0].status.conditions[-1].type}={.items[0].status.conditions[-1].status}").To(status)
		if status.String() == "Ready=True" {
			break
		}
		time.Sleep(2 * time.Second)
	}
}
