export GO111MODULE=on

.PHONY: all
all: deps zouch

.PHONY: deps
deps:
	go get -d -v
	go mod tidy

zouch: $(shell find . -name "*.go")
	go build -v

.PHONY: test
test: deps
	go test -v ./...

.PHONY: clean
clean:
	${RM} zouch
