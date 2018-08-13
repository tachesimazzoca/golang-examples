GO_EXECUTABLE ?= go
GOPATH := $(CURDIR):$(GOPATH)

.PHONY: install-deps
install-deps:
	@rm -rf src/hello-godep/Godeps
	@rm -rf src/hello-godep/vendor
	@go get -u "github.com/mitchellh/cli"
	@go get -u "github.com/mattn/go-colorable"
	@cd src/hello-godep && godep save
	@rm -rf src/github.com
	@rm -rf src/golang.org

.PHONY: install
install:
	@$(GO_EXECUTABLE) install hello-godep

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf src/hello-godep/Godeps
	@rm -rf src/hello-godep/vendor

