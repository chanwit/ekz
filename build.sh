TAG=v1.18.9-eks-1-18-1.1

mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build v0.8.1

git apply ../../patches/001_dockerfile.patch
git apply ../../patches/002_embedded-bins_Makefile.patch
git apply ../../patches/003_pkg_constant.patch

make EMBEDDED_BINS_BUILDMODE=fetch
docker login quay.io -u $QUERY_USERNAME -p $QUERY_PASSWORD
docker build -t quay.io/chanwit/ekz:$TAG .
docker push quay.io/chanwit/ekz:$TAG
docker logout
