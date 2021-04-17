BUILD="0"
EKSD_YAML="eksd.yaml"

wget -qO- https://distro.eks.amazonaws.com/kubernetes-1-19/kubernetes-1-19-eks-3.yaml > ${EKSD_YAML}

CHANNEL=$(yq eval '.spec.channel' $EKSD_YAML)
NUMBER=$(yq eval '.spec.number' $EKSD_YAML)
VERSION=$(yq e '.status.components[] | select(.name=="kubernetes").gitTag | split("v") | .[1]' $EKSD_YAML)


TAG="v${VERSION}-eks-${CHANNEL}-${NUMBER}.${BUILD}"
echo $TAG > TAG

docker build \
  --build-arg K8S_VERSION=v${VERSION} \
  --build-arg EKSD_CHANNEL=$CHANNEL \
  --build-arg EKSD_NUMBER=$NUMBER \
  -t quay.io/ekz-io/kind:$TAG .
