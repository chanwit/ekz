K9S_VERSION=$(cat K9S_VERSION)

rm -rf build/

mkdir build/
cd build/
git clone https://github.com/derailed/k9s
cd k9s
git checkout -b build $K9S_VERSION
stg init
stg import -t -s ../../patches/series

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
