package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a cluster",
	Long:  "The create sub-commands create EKS-D clusters.",
	RunE:  createClusterCmdRun,
}

var (
	provider       string
	eksdVersion    string
	kubeConfigFile string
)

func init() {
	createClusterCmd.Flags().StringVar(&provider, "provider", "ekz", "cluster provider (ekz, kind)")
	createClusterCmd.Flags().StringVar(&eksdVersion, "eksd-version", "v1.18.9-eks-1-18-1", "specify a version of EKS-D")
	createClusterCmd.Flags().StringVarP(&kubeConfigFile, "output", "o", "kubeconfig", "specify output file to write kubeconfig to")

	createCmd.AddCommand(createClusterCmd)
}

func createClusterCmdRun(cmd *cobra.Command, args []string) error {
	switch provider {
	case "ekz":
		return createClusterEKZRun(cmd, args)
	case "kind":
		return createClusterKINDRun(cmd, args)
	default:
		return fmt.Errorf("NYI provider: %s", provider)
	}
}

func createClusterEKZRun(cmd *cobra.Command, args []string) error {
	// "3" is the latest stable provided by EKZ
	ekzImageBuild := "3"
	imageName := fmt.Sprintf("quay.io/chanwit/ekz:%s.%s", eksdVersion, ekzImageBuild)
	containerName := "ekz-controller-0"

	logger.Actionf("pulling image: %s ...", imageName)
	var err error
	err = script.Exec("docker", "pull", imageName).Run()
	if err != nil {
		return errors.Wrapf(err, "error pulling image: %s", imageName)
	}

	containerId := script.Var()
	err = script.Exec("docker", "ps", "-aq", "-f", "name="+containerName).To(containerId)
	if err != nil {
		return errors.Wrapf(err, "failed to run docker ps to check container: %s.", containerName)
	}

	// container existed
	if containerId.String() != "" {
		return errors.Wrapf(err, "container %s existed. cluster creation aborted.", containerName)
	}

	logger.Actionf("starting container: %s ...", containerName)
	_, stderr, err := script.Exec("docker", "run",
		"--detach",
		"--name", containerName,
		"--hostname", "controller",
		"--privileged",
		"-v", "/var/lib/ekz",
		"-p", "6443:6443",
		imageName).
		DividedOutput()
	if err != nil {
		return errors.Wrapf(err, "failed to start %s container with image: %s. %s", containerName, imageName, strings.TrimSpace(string(stderr)))
	}
	// handle port clash
	// handle container name clash

	// TODO use retry-backoff
	time.Sleep(2 * time.Second)

	// TODO append kubeconfig to ~/.kube/config
	err = script.Exec("docker", "exec",
		containerName,
		"cat", "/var/lib/ekz/pki/admin.conf").
		WriteFile(kubeConfigFile, 0644).
		Run()
	if err != nil {
		return errors.Wrapf(err, "error obtaining kubeconfig from container: %s", containerName)
	}

	logger.Waitingf("an EKS-D cluster is now being provisioned ...")
	// TODO wait until node ready

	logger.Waitingf("waiting for cluster to start ...")
	for {
		name := script.Var()
		script.Exec("kubectl", "--kubeconfig="+kubeConfigFile, "get", "nodes", "-ojsonpath={.items[0].metadata.name}").To(name)
		if name.String() == "controller" {
			break
		}
		time.Sleep(2 * time.Second)
	}

	logger.Waitingf("waiting for cluster to be ready ...")
	for {
		status := script.Var()
		script.Exec("kubectl", "--kubeconfig="+kubeConfigFile, "get", "nodes", "-ojsonpath={.items[0].status.conditions[-1].type}={.items[0].status.conditions[-1].status}").To(status)
		if status.String() == "Ready=True" {
			break
		}
		time.Sleep(2 * time.Second)
	}

	logger.Successf("an EKS-D cluster is now ready.")
	return nil
}

func createClusterKINDRun(cmd *cobra.Command, args []string) error {
	return nil
}
