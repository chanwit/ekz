FROM kindest/node:v1.19.7@sha256:a70639454e97a4b733f9d9b67e12c01f6b0297449d5b9cbbef87473458e26dca

ARG K8S_VERSION=v1.19.6-eks-1-19-1
ARG BASE_URL=https://distro.eks.amazonaws.com/kubernetes-1-19/releases/1/artifacts/kubernetes/v1.19.6/bin/linux/amd64

ADD $BASE_URL/kubectl /usr/bin/kubectl
ADD $BASE_URL/kubeadm /usr/bin/kubeadm
ADD $BASE_URL/kubelet /usr/bin/kubelet
RUN echo $K8S_VERSION > /kind/version
RUN chmod +x /usr/bin/kubectl /usr/bin/kubeadm /usr/bin/kubelet
