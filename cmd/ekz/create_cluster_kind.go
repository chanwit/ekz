package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"sigs.k8s.io/kind/pkg/cluster"
)

func createClusterKIND() error {
	buildNo := "1" // KIND v0.10.0 with 1.18.15 node image

	image := fmt.Sprintf("quay.io/ekz-io/kind:%s.%s", eksdVersion, buildNo)

	logger.Actionf("pulling image %s ...", image)
	err := script.Exec("docker", "pull", image).Run()
	if err != nil {
		return err
	}

	provider := cluster.NewProvider()
	os.Setenv("KIND_EXPERIMENTAL_DOCKER_NETWORK", fmt.Sprintf("%s-bridge", clusterName))

	parts := strings.SplitN(eksdVersion, "-", 2)
	if len(parts) != 2 {
		return errors.Errorf("eksdVersion: %s cannot be split into two", eksdVersion)
	}

	suffix := parts[1]
	config := fmt.Sprintf(`
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  disableDefaultCNI: true   # disable kindnet
  podSubnet: 192.168.0.0/16 # set to Calico's default subnet
kubeadmConfigPatches:
- |
  apiVersion: kubeadm.k8s.io/v1beta2
  kind: ClusterConfiguration
  metadata:
    name: config
  imageRepository: "public.ecr.aws/eks-distro/kubernetes"
  kubernetesVersion: "%s"
  etcd:
    local:
      imageRepository: "public.ecr.aws/eks-distro/etcd-io"
      imageTag: "v3.4.14-%s"
  dns:
    imageRepository: "public.ecr.aws/eks-distro/coredns"
    imageTag: "v1.7.0-%s"
  apiServer:
    extraArgs:
      service-account-issuer: "kubernetes.default.svc"
      service-account-signing-key-file: "/etc/kubernetes/pki/sa.key"
`, eksdVersion, suffix, suffix)

	logger.Actionf("creating cluster: %s ...", clusterName)

	if err := provider.Create(
		clusterName,
		cluster.CreateWithRawConfig([]byte(config)),
		cluster.CreateWithNodeImage(image),
		cluster.CreateWithRetain(false),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	); err != nil {
		return errors.Wrapf(err, "failed to create cluster %v", clusterName)
	}

	if err := provider.ExportKubeConfig(clusterName, kubeConfigFile); err != nil {
		return err
	}
	logger.Successf("kubeconfig is written to: %s", kubeConfigFile)

	logger.Waitingf("waiting for cluster to start ...")
	waitForNodeStarted(fmt.Sprintf("%s-control-plane", clusterName), 30*time.Second)

	// apply Calico
	logger.Waitingf("applying Calico manifests ...")
	if err := script.Exec("kubectl", "apply", "-f", "https://docs.projectcalico.org/v3.16/manifests/calico.yaml").Run(); err != nil {
		// Relax Calico's RPF Check Configuration
		// By default, Calico pods fail if the Kernel's Reverse Path Filtering (RPF) check is not enforced.
		// This is a security measure to prevent endpoints from spoofing their IP address.
		// The RPF check is not enforced in Kind nodes.
		// Thus, we need to disable the Calico check by setting an environment variable in the calico-node DaemonSet:
		if err := script.Exec("kubectl", "-n", "kube-system", "set", "env", "daemonset/calico-node", "FELIX_IGNORELOOSERPF=true").Run(); err != nil {
			return errors.Wrapf(err, "could not able to relex RPF")
		}
		return errors.Wrapf(err, "could not apply Calico manifests")
	}

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady(60 * time.Second)

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}
