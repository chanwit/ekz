VERSION?=$(shell grep 'VERSION' cmd/flux/main.go | awk '{ print $$4 }' | tr -d '"')

all: test build

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

test: tidy fmt vet docs
	go test ./... -coverprofile cover.out

build-ui:
	( cd ui && bash -x ./dev.sh )
	( cd ui && bash -x ./build.sh )

build:
	CGO_ENABLED=0 go build -o ./bin/ekz ./cmd/ekz

install:
	go install cmd/ekz

.PHONY: docs
docs:
	rm docs/cmd/ekz* || echo "failure!"
	mkdir -p ./docs/cmd && go run ./cmd/ekz/ docgen

install-dev: build
	sudo cp ./bin/ekz /usr/local/bin