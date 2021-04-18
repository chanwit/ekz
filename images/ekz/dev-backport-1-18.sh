TAG=$(cat VERSION_1_18)
K0S_VERSION=$(cat K0S_VERSION)

mkdir build/
cd build/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build $K0S_VERSION
stg init
stg import -t -s ../../patches-backport-1-18/series

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
