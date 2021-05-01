# at "ui" dir
(cd build/k9s && make build)

docker build -t quay.io/ekz-io/ekz-webui:latest \
       -f Dockerfile.webui .
