package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func createClusterEKZ() error {
	// "3" is the latest stable provided by EKZ
	ekzImageBuild := "3"
	imageName := fmt.Sprintf("quay.io/ekz-io/ekz:%s.%s", eksdVersion, ekzImageBuild)
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
		"--label", "io.x-k8s.ekz.cluster=ekz",
		"-v", "/var/lib/ekz",
		"-p", "6443:6443",
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

	logger.Waitingf("an EKS-D cluster is now being provisioned ...")

	logger.Waitingf("waiting for cluster to start ...")
	waitForNodeStarted("controller")

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady()

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}
