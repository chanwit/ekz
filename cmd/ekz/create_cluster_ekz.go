package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
)

func createClusterEKZ() error {
	// "4" is the latest stable provided by EKZ
	ekzImageBuild := "4"
	imageName := fmt.Sprintf("quay.io/ekz-io/ekz:%s.%s", eksdVersion, ekzImageBuild)
	containerName := fmt.Sprintf("%s-controller-0", clusterName)

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
		if verbose {
			fmt.Printf("[DEBUG] containerId.String()=%s\n", containerId.String())
		}

		return errors.Errorf("container %s existed - cluster creation aborted", containerName)
	}

	logger.Actionf("starting container: %s ...", containerName)
	_, stderr, err := script.Exec("docker", "run",
		"--detach",
		"--name", containerName,
		"--hostname", "controller",
		"--privileged",
		"--label", fmt.Sprintf("io.x-k8s.ekz.cluster=%s", clusterName),
		"-v", "/var/lib/ekz",
		"-p", "127.0.0.1:0:6443",
		imageName).
		DividedOutput()
	if err != nil {
		return errors.Wrapf(err, "failed to start %s container with image: %s. %s", containerName, imageName, strings.TrimSpace(string(stderr)))
	}

	// TODO use retry-backoff instead of fixing 2 seconds here
	time.Sleep(2 * time.Second)

	// TODO handle port clash
	// TODO handle container name clash
	err = getKubeconfigEKZ(containerName, kubeConfigFile)
	if err != nil {
		return err
	}
	logger.Successf("kubeconfig is written to: %s", kubeConfigFile)

	logger.Waitingf("waiting for cluster to start ...")
	waitForNodeStarted("controller")

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady()

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}
