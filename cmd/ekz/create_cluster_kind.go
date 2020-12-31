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

	image := fmt.Sprintf("quay.io/ekz-io/kind:%s", eksdVersion)

	logger.Actionf("pulling image %s ...", image)
	err := script.Exec("docker", "pull", image).Run()
	if err != nil {
		return err
	}

	provider := cluster.NewProvider()
	os.Setenv("KIND_EXPERIMENTAL_DOCKER_NETWORK", fmt.Sprintf("%s-bridge", clusterName))

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

	logger.Waitingf("waiting for cluster to be ready ...")
	waitForNodeReady(60 * time.Second)

	logger.Successf("the EKS-D cluster is now ready.")
	return nil
}
