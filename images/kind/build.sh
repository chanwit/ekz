case $1 in
  1.18)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.18.15@sha256:5c1b980c4d0e0e8e7eb9f36f7df525d079a96169c8a8f20d8bd108c0d0889cc4"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-18/kubernetes-1-18-eks-3.yaml"
    ;;

  1.19)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.19.7@sha256:a70639454e97a4b733f9d9b67e12c01f6b0297449d5b9cbbef87473458e26dca"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-19/kubernetes-1-19-eks-3.yaml"
    ;;
esac

EKSD_YAML="eksd.yaml"
wget -qO- $URL > ${EKSD_YAML}
CHANNEL=$(yq eval '.spec.channel' $EKSD_YAML)
NUMBER=$(yq eval '.spec.number' $EKSD_YAML)
VERSION=$(yq e '.status.components[] | select(.name=="kubernetes").gitTag | split("v") | .[1]' $EKSD_YAML)
TAG="v${VERSION}-eks-${CHANNEL}-${NUMBER}.${BUILD}"
echo $TAG > TAG

docker build \
  --build-arg BASE_IMAGE=$BASE_IMAGE \
  --build-arg K8S_VERSION=v${VERSION} \
  --build-arg EKSD_CHANNEL=$CHANNEL \
  --build-arg EKSD_NUMBER=$NUMBER \
  --build-arg EKSD_BASE_URL=https://distro.eks.amazonaws.com/kubernetes-${$CHANNEL}/releases/${NUMBER}/artifacts/kubernetes/v${VERSION}/bin/linux/amd64
  -t quay.io/ekz-io/kind:$TAG .
