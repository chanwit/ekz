#
# args 1 is one of 1-18, 1-19, 1-20 or 1-21
#
RELEASE=$1

REL="${RELEASE//-/_}"

TAG=$(cat VERSION_${REL})
K0S_VERSION=$(cat K0S_VERSION)

mkdir build-${RELEASE}/
cd build-${RELEASE}/
git clone https://github.com/k0sproject/k0s
cd k0s
git checkout -b build-${RELEASE} $K0S_VERSION
stg init
stg import -t -s ../../patches-${RELEASE}/series

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