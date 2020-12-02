mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build v0.8.0

git apply ../../patches/001_dockerfile.patch
git apply ../../patches/002_embedded-bins_Makefile.patch
git apply ../../patches/003_pkg_constant.patch

make EMBEDDED_BINS_BUILDMODE=fetch
docker build -t chanwit/ekz:v1.18.9-eks-1-18-1 .
docker push chanwit/ekz:v1.18.9-eks-1-18-1

# docker exec ekz-controller cat /var/lib/ekz/pki/admin.conf > ~/.kube/config

# docker run --name ekz-controller \
#    --hostname controller \
#    --privileged -v /var/lib/ekz \
#    -p 6443:6443 chanwit/ekz:v1.18.9-eks-1-18-1