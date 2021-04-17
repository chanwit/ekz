BUILD="0"

EKSD_YAML="eksd.yaml"

wget -O- https://distro.eks.amazonaws.com/kubernetes-1-19/kubernetes-1-19-eks-3.yaml > ${EKSD_YAML}

CHANNEL=$(yq eval '.spec.channel' $EKSD_YAML)
NUMBER=$(yq eval '.spec.number' $EKSD_YAML)

VERSION=$(yq e '.status.components[] | select(.name=="kubernetes").gitTag | split("v") | .[1]' $EKSD_YAML)

# echo to VERSION file
echo "v${VERSION}-eks-${CHANNEL}-${NUMBER}.${BUILD}" > VERSION

