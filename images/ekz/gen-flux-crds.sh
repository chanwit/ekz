#!/usr/bin/env bash

version=${1:-master}
# echo "Flux version: $version"

# echo "Building the CLI image ..."
docker build -t flux-cli:v${version} \
  --build-arg FLUXCLI_VER=${version} \
  -f Dockerfile.fluxcli . > /dev/null

#
# echo "Generating the manifests using the built CLI ..."
#
# excludes: services, deployment
#
manifest="manifests-$version.yaml"
docker run --rm -it flux-cli:v${version} install --version="$version" \
  --components-extra=image-reflector-controller,image-automation-controller \
  --export --dry-run \
| yq e '. | select(.kind != "Service" and .kind != "Deployment")' - \
> ./build/k0s/static/manifests/flux/CustomResourceDefinition/flux-system.yaml

