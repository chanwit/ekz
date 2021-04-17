FROM kindest/node:v1.19.7@sha256:a70639454e97a4b733f9d9b67e12c01f6b0297449d5b9cbbef87473458e26dca

ARG K8S_VERSION=v1.19.6
ARG EKSD_CHANNEL=1-19
ARG EKSD_NUMBER=3
ARG EKSD_VERSION=${K8S_VERSION}-eks-${EKSD_CHANNEL}-${EKSD_NUMBER}

ARG BASE_URL=https://distro.eks.amazonaws.com/kubernetes-${EKSD_CHANNEL}/releases/${EKSD_NUMBER}/artifacts/kubernetes/${K8S_VERSION}/bin/linux/amd64

ADD $BASE_URL/kubectl /usr/bin/kubectl
ADD $BASE_URL/kubeadm /usr/bin/kubeadm
ADD $BASE_URL/kubelet /usr/bin/kubelet
RUN echo ${EKSD_VERSION} > /kind/version
RUN chmod +x /usr/bin/kubectl /usr/bin/kubeadm /usr/bin/kubelet
