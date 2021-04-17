ARG BASE_IMAGE

FROM $BASE_IMAGE

ARG K8S_VERSION
ARG EKSD_CHANNEL
ARG EKSD_NUMBER
ARG EKSD_BASE_URL
ARG EKSD_VERSION

ADD $EKSD_BASE_URL/kubectl /usr/bin/kubectl
ADD $EKSD_BASE_URL/kubeadm /usr/bin/kubeadm
ADD $EKSD_BASE_URL/kubelet /usr/bin/kubelet
RUN echo ${EKSD_VERSION} > /kind/version
RUN chmod +x /usr/bin/kubectl /usr/bin/kubeadm /usr/bin/kubelet
