package main

import (
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kind/pkg/cluster"
)

var getKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Get kubeconfig",
	Long:  "This command obtains the KubeConfig of the EKS-D cluster and writes to the target file.",
	Example: `  # Get the KubeConfig from the cluster and write to $PWD/kubeconfig
  ekz get kubeconfig

  # Get the KubeConfig and writes to $HOME/.kube/config
  # Please note that this example overwrites the content of $HOME/.kube/config file
  ekz get kubeconfig -o $HOME/.kube/config
`,
	RunE: getKubeconfigCmdRun,
}

func init() {
	getKubeconfigCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", "kubeconfig", "specify output file to write kubeconfig to")

	getCmd.AddCommand(getKubeconfigCmd)
}

func getKubeconfigCmdRun(cmd *cobra.Command, args []string) error {
	switch provider {
	case "ekz":
		return getKubeconfigEKZ("ekz-controller-0", kubeConfigFile)
	case "kind":
		return getKubeconfigKIND()
	}

	return nil
}

func getKubeconfigEKZ(containerName string, targetFile string) error {
	// TODO append kubeconfig to ~/.kube/config
	err := script.Exec("docker", "exec",
		containerName,
		"cat", "/var/lib/ekz/pki/admin.conf").
		WriteFile(targetFile, 0644).
		Run()
	if err != nil {
		return errors.Wrapf(err, "error obtaining kubeconfig from container: %s", containerName)
	}
	return nil
}

func getKubeconfigKIND() error {
	clusterName := "ekz"
	provider := cluster.NewProvider()
	return provider.ExportKubeConfig(clusterName, kubeConfigFile)
}
