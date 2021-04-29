export GO111MODULE=on

VERSION := $(shell git describe --tags --abbrev=0)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date "+%Y-%m-%d %H:%M:%S")
LDFLAGS := -X "main.version=${VERSION}"
LDFLAGS += -X "main.commit=${COMMIT}"
LDFLAGS += -X "main.date=${DATE}"

.PHONY: all
all: deps zouch

.PHONY: deps
deps:
	go get -d -v
	go mod tidy

zouch: $(shell find . -name "*.go")
	go build -v --ldflags='${LDFLAGS}'

.PHONY: test
test: deps
	go test -v ./...

.PHONY: lint
lint: 
	golangci-lint run -v ./...

.PHONY: clean
clean:
	${RM} zouch
