VERSION := $(shell cat VERSION)

deps:
	glide -q install

test: deps
	go test `glide nv`

build: test
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags production -o pkg/linux/amd64/go-shippo.a
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -tags production -o pkg/darwin/amd64/go-shippo.a

.PHONY: build test deps