GO_EXECUTABLE ?= go
GOPATH := $(CURDIR):$(GOPATH)

.PHONY: install-deps
install-deps:
	@cd src/hello-glide && glide update

.PHONY: run
run:
	@$(GO_EXECUTABLE) run src/hello-glide/main.go

.PHONY: install
install:
	@$(GO_EXECUTABLE) install hello-glide

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf src/hello-glide/vendor
	@rm -rf src/hello-glide/glide.lock
