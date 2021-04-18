package main

import (
	"fmt"
	"time"

	"github.com/chanwit/ekz/pkg/constants"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a cluster",
	Args:  cobra.MaximumNArgs(1),
	Long:  "The create sub-commands create EKS-D clusters.",
	Example: `  # Create an EKS-D cluster with the default provider
  # The KubeConfig will be merged to $HOME/.kube/config. The default cluster name is 'ekz'.
  ekz create cluster

  # Create cluster and name it 'dev'
  ekz create cluster --name=dev

  # Create the 'dev' cluster (alternative syntax)
  ekz create cluster dev

  # Create the default cluster in the host mode
  # This command runs the cluster using all (net,ipc,pid,uts) host namespaces, 
  # similar to run it directly on the local machine. 
  ekz create cluster --host

  # Create an EKS-D cluster with the EKZ provider
  # This command creates an EKS-D-compatible K0s-based cluster.
  ekz --provider=ekz create cluster

  # Create an EKS-D cluster with the KIND provider
  # This command creates an EKS-D-compatible KIND cluster.
  ekz --provider=kind create cluster

  # Create an EKS-D cluster and write KubeConfig to $PWD/kubeconfig
  # If the file already exists, the new KubeConfig will be merged into it.
  ekz create cluster -o kubeconfig

  # Create EKS-D cluster with a specific version of EKS-D
  ekz create --eksd-version=v1.18.9-eks-1-18-1 cluster

  # Create EKS-D cluster with a short version format
  # Please use v1.18 for v1.18.9-eks-1-18-3, and v1.19 for v1.19.6-eks-1-19-3.
  ekz create --eksd-version=v1.18 cluster
`,
	RunE: createClusterCmdRun,
}

var (
	eksdVersion           string
	kubeConfigFile        string
	clusterName           string
	hostMode              bool
	hostModeVolumeMapping bool
)

func init() {
	createClusterCmd.Flags().StringVar(&eksdVersion, "eksd-version", "v1.19.6-eks-1-19-3", "specify a version of EKS-D")
	createClusterCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", constants.BackTickHomeFile, "specify output file to write kubeconfig to")
	createClusterCmd.Flags().StringVar(&clusterName, "name", "ekz", "cluster name")
	createClusterCmd.Flags().BoolVar(&hostMode, "host", false, "run in the host mode")
	createClusterCmd.Flags().BoolVar(&hostModeVolumeMapping, "map-to-host-volume", false, "map /var/lib/k0s to the host directory")

	createCmd.AddCommand(createClusterCmd)
}

func createClusterCmdRun(cmd *cobra.Command, args []string) error {
	logger.Successf("the default provider is: %s", provider)

	// use args[0] as the clusterName
	if len(args) == 1 {
		clusterName = args[0]
	}

	if eksdVersion == "v1.18" {
		eksdVersion = "v1.18.9-eks-1-18-3"
	} else if eksdVersion == "v1.19" {
		eksdVersion = "v1.19.6-eks-1-19-3"
	}

	// TODO validate eksdVersion
	// v1.18.9-eks-1-18-1
	// v1.19.6-eks-1-19-1

	switch provider {
	case "ekz":
		return createClusterEKZ()
	case "kind":
		if hostMode == true {
			return errors.New("the host mode is not supported by the KIND provider")
		}
		return createClusterKIND()
	default:
		return fmt.Errorf("NYI provider: %s", provider)
	}
}

func waitForNodeStarted(nodeName string, timeout time.Duration) {
	for start := time.Now(); time.Since(start) < timeout; {
		name := script.Var()
		script.Exec("kubectl", "--kubeconfig="+expandKubeConfigFile(), "get", "nodes", "-ojsonpath={.items[0].metadata.name}").To(name)
		if name.String() == nodeName {
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func waitForNodeReady(timeout time.Duration) {
	for start := time.Now(); time.Since(start) < timeout; {
		status := script.Var()
		script.Exec("kubectl", "--kubeconfig="+expandKubeConfigFile(),
			"get", "nodes",
			"-ojsonpath={.items[0].status.conditions[-1].type}={.items[0].status.conditions[-1].status}").To(status)
		if status.String() == "Ready=True" {
			break
		}
		time.Sleep(10 * time.Second)
	}
}
