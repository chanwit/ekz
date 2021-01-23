FROM kindest/node:v1.18.15@sha256:5c1b980c4d0e0e8e7eb9f36f7df525d079a96169c8a8f20d8bd108c0d0889cc4

ARG K8S_VERSION=v1.18.9-eks-1-18-1
ARG BASE_URL=https://distro.eks.amazonaws.com/kubernetes-1-18/releases/1/artifacts/kubernetes/v1.18.9/bin/linux/amd64

ADD $BASE_URL/kubectl /usr/bin/kubectl
ADD $BASE_URL/kubeadm /usr/bin/kubeadm
ADD $BASE_URL/kubelet /usr/bin/kubelet
RUN echo $K8S_VERSION > /kind/version
RUN chmod +x /usr/bin/kubectl /usr/bin/kubeadm /usr/bin/kubelet
