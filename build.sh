TAG=$(cat VERSION)

mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build v0.8.1
stg import -s ../../patches/series

make EMBEDDED_BINS_BUILDMODE=fetch
docker build -t quay.io/chanwit/ekz:$TAG .
