package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"os"
	"sigs.k8s.io/kind/pkg/cluster"
	"strings"
)

func createClusterKIND() error {

	image := fmt.Sprintf("quay.io/ekz-io/kind:%s", eksdVersion)

	logger.Actionf("pulling image %s ...", image)
	err := script.Run("docker", "pull", image)
	if err != nil {
		return err
	}
	name := "ekz"

	provider := cluster.NewProvider()
	os.Setenv("KIND_EXPERIMENTAL_DOCKER_NETWORK", "ekz-bridge")

	parts := strings.SplitN(eksdVersion, "-", 2)
	suffix := parts[1]
	config := fmt.Sprintf(`
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
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
`, eksdVersion, suffix, suffix)

	logger.Actionf("creating cluster: %s ...", name)

	if err := provider.Create(
		name,
		cluster.CreateWithRawConfig([]byte(config)),
		cluster.CreateWithNodeImage(image),
		cluster.CreateWithRetain(false),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	); err != nil {
		return errors.Wrapf(err, "failed to create cluster %v", name)
	}

	if err := provider.ExportKubeConfig(name, kubeConfigFile); err != nil {
		return err
	}
	logger.Successf("kubeconfig is written to: %s", kubeConfigFile)

	logger.Waitingf("waiting for cluster to start ...")
	waitForNodeStarted("ekz-control-plane")

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady()

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}
