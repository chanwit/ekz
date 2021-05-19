package main

import "fmt"

func getKindConfig(eksdVersion string, suffix string) (string, error) {
	switch eksdVersion {
	case "v1.18.16-eks-1-18-4":
		return getKindConfig1_18(eksdVersion, suffix), nil
	case "v1.19.8-eks-1-19-4":
		return getKindConfig1_19(eksdVersion, suffix), nil
	case "v1.20.4-eks-1-20-1":
		return getKindConfig1_20(eksdVersion, suffix), nil
	}
	// TODO return error
	return "", fmt.Errorf("getKindConfig: does not support eksVersion: %s", eksdVersion)
}

func getKindConfig1_18(eksdVersion string, suffix string) string {
	return fmt.Sprintf(`
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
    - hostPath: /var/run/docker.sock
      containerPath: /var/run/docker.sock
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
}

func getKindConfig1_19(eksdVersion string, suffix string) string {
	return fmt.Sprintf(`
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
    - hostPath: /var/run/docker.sock
      containerPath: /var/run/docker.sock
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
    imageTag: "v1.8.0-%s"
  apiServer:
    extraArgs:
      service-account-issuer: "kubernetes.default.svc"
      service-account-signing-key-file: "/etc/kubernetes/pki/sa.key"
`, eksdVersion, suffix, suffix)
}

func getKindConfig1_20(eksdVersion string, suffix string) string {
	return fmt.Sprintf(`
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
    - hostPath: /var/run/docker.sock
      containerPath: /var/run/docker.sock
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
      imageTag: "v3.4.15-%s"
  dns:
    imageRepository: "public.ecr.aws/eks-distro/coredns"
    imageTag: "v1.8.3-%s"
  apiServer:
    extraArgs:
      service-account-issuer: "kubernetes.default.svc"
      service-account-signing-key-file: "/etc/kubernetes/pki/sa.key"
`, eksdVersion, suffix, suffix)
}
