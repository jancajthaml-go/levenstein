.PHONY: all
all: build test benchmark

.PHONY: build
build:
	@go build levenstein.go main.go

.PHONY: test
test:
	@go test

.PHONY: benchmark
benchmark:
	@go test -bench=.