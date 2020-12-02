mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build v0.8.0

git apply ../../patches/001_dockerfile.patch
git apply ../../patches/002_embedded-bins_Makefile.patch
git apply ../../patches/003_pkg_constant.patch

make EMBEDDED_BINS_BUILDMODE=fetch
docker build -t docker.pkg.github.com/chanwit/ekz/ekz:v1.18.9-eks-1-18-1 .
docker push docker.pkg.github.com/chanwit/ekz/ekz:v1.18.9-eks-1-18-1
