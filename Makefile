VERSION ?= $(shell git describe --tags --always --dirty)

.PHONY: test

test:
	cd http && \
	go test -count=1 -v ./... && \
	cd ../fixtures/db && \
	go test -count=1 -v ./...
