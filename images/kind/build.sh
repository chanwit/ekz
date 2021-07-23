case $1 in
  1.18)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.18.19@sha256:7af1492e19b3192a79f606e43c35fb741e520d195f96399284515f077b3b622c"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-18/kubernetes-1-18-eks-7.yaml"
    ;;

  1.19)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.19.11@sha256:07db187ae84b4b7de440a73886f008cf903fcf5764ba8106a9fd5243d6f32729"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-19/kubernetes-1-19-eks-6.yaml"
    ;;

  1.20)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.20.7@sha256:cbeaf907fc78ac97ce7b625e4bf0de16e3ea725daf6b04f930bd14c67c671ff9"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-20/kubernetes-1-20-eks-3.yaml"
    ;;

  1.21)
    BUILD="0"
    BASE_IMAGE="kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6"
    URL="https://distro.eks.amazonaws.com/kubernetes-1-21/kubernetes-1-21-eks-1.yaml"
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
  --build-arg EKSD_VERSION=v${VERSION}-eks-${CHANNEL}-${NUMBER} \
  --build-arg EKSD_BASE_URL=https://distro.eks.amazonaws.com/kubernetes-${CHANNEL}/releases/${NUMBER}/artifacts/kubernetes/v${VERSION}/bin/linux/amd64 \
  -t quay.io/ekz-io/kind:$TAG .
