TAG=$(cat VERSION_1_18)
K0S_VERSION=$(cat K0S_VERSION)

mkdir build-1-18/
cd build-1-18/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build $K0S_VERSION
stg init
stg import -t -s ../../patches-1-18/series

# Workflow
# ========

# hack
#   git commit
#   stg repair
# or
# hack
#   stg add
#   stg refresh
# or
# hack
#   commit
#   stg uncommit
# then
#   stg export -n -p -d ../../patches/
