GOTEST_FLAGS=-cpu=1,2,4

APP=lightaws
MAIN_FILE=main.go

BASE_PACKAGE=github.com/openfresh/lightaws
deps:
		go get github.com/golang/dep
		go get github.com/golang/lint/golint
		dep ensure

define build-artifact
		GOOS=$(1) GOARCH=$(2) go build -o artifacts/$(APP)
		cd artifacts && tar cvzf $(APP)_$(1)_$(2).tar.gz $(APP)
		rm ./artifacts/$(APP)
		@echo [INFO]build success: $(1)_$(2)
endef

build-all:
		$(call build-artifact,linux,386)
		$(call build-artifact,linux,amd64)
		$(call build-artifact,linux,arm)
		$(call build-artifact,linux,arm64)
		$(call build-artifact,darwin,amd64)
		$(call build-artifact,windows,386)
		$(call build-artifact,windows,amd64)

build:
		go build -o bin/$(APP) $(MAIN_FILE)

test: vet
		go test -cover $(TARGETS)

vet:
		go vet $(TARGETS)
