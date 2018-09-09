.PHONY: install-deps
install-deps:
	@dep ensure

.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@mkdir -p build
	@go build -o build/hello-database

.PHONY: clean
clean:
	@rm -rf build
	@rm -rf vendor
	@rm -rf Gopkg.lock
