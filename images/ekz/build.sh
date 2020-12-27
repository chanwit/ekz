TAG=$(cat VERSION)
K0S_VERSION=$(cat K0S_VERSION)

mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build $K0S_VERSION
stg init
stg import -s ../../patches/series

make EMBEDDED_BINS_BUILDMODE=fetch
docker build -t quay.io/chanwit/ekz:$TAG .
