GOTEST_FLAGS=-cpu=1,2,4

default: deps

BASE_PACKAGE=github.com/openfresh/lightaws

deps:
		go get github.com/golang/dep
		go get github.com/golang/lint/golint
		dep ensure

build:
		go build -o bin/lightaws main.go

test: vet
		go test -cover $(TARGETS)

vet:
		go vet $(TARGETS)
