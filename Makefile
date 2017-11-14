.PHONY: build default

default: build

build:
	go build -o goland .
	cp -r goland "${GOPATH}/bin"


