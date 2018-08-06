GO_EXECUTABLE ?= go
GOPATH := $(CURDIR):$(GOPATH)

.PHONY: install-deps
install-deps:
	@cd src/hello-dep && dep ensure

.PHONY: run
run:
	@$(GO_EXECUTABLE) run src/hello-dep/main.go

.PHONY: install
install:
	@$(GO_EXECUTABLE) install hello-dep

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf src/hello-dep/vendor
