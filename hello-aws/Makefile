.PHONY: install-deps
install-deps:
	@dep ensure

.PHONY: build
build:
	@mkdir -p build
	@go build -o build/hello-aws

.PHONY: clean
clean:
	@rm -rf build
	@rm -rf vendor
	@rm -rf Gopkg.lock
