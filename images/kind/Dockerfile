FROM kindest/node:v1.18.8

ARG K8S_VERSION=v1.18.9-eks-1-18-1
ARG BASE_URL=https://distro.eks.amazonaws.com/kubernetes-1-18/releases/1/artifacts/kubernetes/v1.18.9/bin/linux/amd64

ADD $BASE_URL/kubectl /usr/bin/kubectl
ADD $BASE_URL/kubeadm /usr/bin/kubeadm
ADD $BASE_URL/kubelet /usr/bin/kubelet
RUN echo $K8S_VERSION > /kind/version
RUN chmod +x /usr/bin/kubectl /usr/bin/kubeadm /usr/bin/kubelet
