package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/chanwit/ekz/pkg/constants"
	"github.com/chanwit/ekz/pkg/manifests"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
)

func getEKZNetworkName(containerId string) (string, error) {
	output := script.Var()
	err := script.Exec("docker", "inspect",
		containerId,
		"--format", fmt.Sprintf(`{{index .Config.Labels "%s"}}`, constants.EKZNetworkLabel)).
		To(output)
	if err != nil {
		return "", errors.Wrap(err, "failed to get network name")
	}

	return output.String(), nil
}

func createClusterEKZ() error {
	// build "8" is the latest stable provided by EKZ 1.18
	// build "4" is the latest stable provided by EKZ 1.19
	var ekzImageBuild string
	switch eksdVersion {
	case "v1.18.16-eks-1-18-5":
		ekzImageBuild = "0"
	case "v1.19.8-eks-1-19-4":
		ekzImageBuild = "5"
	case "v1.20.4-eks-1-20-1":
		ekzImageBuild = "2"
	}

	imageName := fmt.Sprintf("quay.io/ekz-io/ekz:%s.%s", eksdVersion, ekzImageBuild)
	containerName := fmt.Sprintf("%s-controller-0", clusterName)

	logger.Actionf("pulling image: %s ...", imageName)
	var err error
	err = script.Exec("docker", "pull", imageName).Run()
	if err != nil {
		return errors.Wrapf(err, "error pulling image: %s", imageName)
	}

	containerId := script.Var()
	err = script.Exec("docker", "ps", "-aq", "-f", fmt.Sprintf("name=%s", containerName)).To(containerId)
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

	if hostMode == false {

		bridgeName := fmt.Sprintf("ekz-%s-bridge", clusterName)

		// TODO check if the bridge already existed

		logger.Actionf("creating bridge network: %s", bridgeName)
		err = script.Exec("docker", "network", "create",
			"-d", "bridge",
			"-o", "com.docker.network.bridge.enable_ip_masquerade=true",
			"-o", "com.docker.network.bridge.enable_icc=true",
			"-o", "com.docker.network.bridge.host_binding_ipv4=0.0.0.0",
			"-o", "com.docker.network.driver.mtu=1500",
			bridgeName).Run()
		if err != nil {
			return errors.Wrapf(err, "failed to create bridge: %s", bridgeName)
		}

		logger.Actionf("starting container: %s ...", containerName)
		_, stderr, err := script.Exec("docker", "run",
			"--detach",
			"--name", containerName,
			"--hostname", "controller",
			"--privileged",
			"--security-opt", "seccomp=unconfined", // also ignore seccomp
			"--security-opt", "apparmor=unconfined", // also ignore apparmor
			// runtime temporary storage
			"--tmpfs", "/tmp", // various things depend on working /tmp
			"--tmpfs", "/run", // systemd wants a writable /run
			"--network", bridgeName,
			"--label", fmt.Sprintf("%s=%s", constants.EKZClusterLabel, clusterName),
			"--label", fmt.Sprintf("%s=%s", constants.EKZNetworkLabel, bridgeName),
			"--volume", "/var/lib/k0s",
			// some k8s things want to read /lib/modules
			"--volume", "/lib/modules:/lib/modules:ro",
			"--volume", "/var/run/docker.sock:/var/run/docker.sock",
			"-p", "127.0.0.1:0:6443",
			imageName).
			DividedOutput()
		if err != nil {
			return errors.Wrapf(err, "failed to start %s container with image: %s. %s", containerName, imageName, strings.TrimSpace(string(stderr)))
		}
	} else if hostMode == true { // MicroK8s-like behavior

		// os.MkdirAll("/var/lib/k0s", 0755)
		volumeMapping := "/var/lib/k0s"
		if hostModeVolumeMapping {
			// persist /var/lib/ekz to host
			volumeMapping = "/var/lib/k0s:/var/lib/k0s"
		}

		logger.Actionf("starting container: %s ...", containerName)
		_, stderr, err := script.Exec("docker", "run",
			"--detach",
			"--name", containerName,
			"--privileged",
			"--security-opt", "seccomp=unconfined", // also ignore seccomp
			"--security-opt", "apparmor=unconfined", // also ignore apparmor
			// runtime temporary storage
			"--tmpfs", "/tmp", // various things depend on working /tmp
			"--volume", "/run:/run", // if host mode we map /run from host
			"--network=host",
			"--ipc=host",
			"--uts=host",
			"--pid=host",
			"--label", fmt.Sprintf("%s=%s", constants.EKZClusterLabel, clusterName),
			"--label", fmt.Sprintf("%s=%s", constants.EKZNetworkLabel, "host"),
			"--volume", volumeMapping,
			// some k8s things want to read /lib/modules
			"--volume", "/lib/modules:/lib/modules:ro",
			"--volume", "/var/run/docker.sock:/var/run/docker.sock",
			imageName).
			DividedOutput()
		if err != nil {
			return errors.Wrapf(err, "failed to start %s container with image: %s. %s", containerName, imageName, strings.TrimSpace(string(stderr)))
		}

	}

	// TODO use retry-backoff instead of fixing 2 seconds here
	time.Sleep(2 * time.Second)

	// TODO handle port clash
	// TODO handle container name clash

	if err := getKubeconfigEKZ(containerName, expandKubeConfigFile()); err != nil {
		return err
	}
	logger.Successf("kubeconfig is written to: %s", expandKubeConfigFile())

	logger.Waitingf("waiting for cluster to start ...")
	waitForNodeStarted("controller", 30*time.Second)

	logger.Actionf("installing the default storage class ...")
	if err := installDefaultStorageClass(); err != nil {
		return err
	}

	if hostMode == false {
		bridgeName := fmt.Sprintf("ekz-%s-bridge", clusterName)
		ipRangeVar := script.Var()
		script.Exec("docker", "inspect", "-f", `{{range .IPAM.Config}}{{.Subnet}}{{end}}`, bridgeName).To(ipRangeVar)
		parts := strings.SplitN(ipRangeVar.String(), ".", 4)
		if len(parts) != 4 {
			return errors.Errorf("error parsing subnet: %s", ipRangeVar.String())
		}

		// TODO: not very a safe range but ok to be .200-220 of a bridge subnet for now
		ipRange := strings.Join(parts[0:3], ".") + ".200-" + strings.Join(parts[0:3], ".") + ".220"

		if verbose {
			logger.Actionf("installing the default load balancer using IP range: %s ...", ipRange)
		} else {
			logger.Actionf("installing the default load balancer ...")
		}

		if err := installLoadBalancer(ipRange); err != nil {
			return err
		}
	}

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady(120 * time.Second)

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}

func installDefaultStorageClass() error {
	return script.Echo(manifests.StorageClass).Exec("kubectl", "--kubeconfig="+expandKubeConfigFile(), "apply", "-f", "-").Run()
}

func installLoadBalancer(ipRange string) error {
	err := script.Echo(manifests.MetalLB).
		Exec("kubectl", "--kubeconfig="+expandKubeConfigFile(), "apply", "-f", "-").
		Run()
	if err != nil {
		return err
	}

	err = script.Echo(fmt.Sprintf(manifests.MetalLBConfig, ipRange)).
		Exec("kubectl", "--kubeconfig="+expandKubeConfigFile(), "apply", "-f", "-").
		Run()
	if err != nil {
		return err
	}

	return nil
}
