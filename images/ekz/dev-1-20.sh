TAG=$(cat VERSION)
K0S_VERSION=$(cat K0S_VERSION)

mkdir build-1-20/
cd build-1-20/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build $K0S_VERSION
stg init
stg import -t -s ../../patches-1-20/series

# Workflow
# ========
# hack
#   stg add
#   stg refresh
# or
# hack
#   commit
#   stg uncommit
# then
#   stg export -n -p -d ../../patches/
