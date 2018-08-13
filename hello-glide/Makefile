.PHONY: install-deps
install-deps:
	@glide update

.PHONY: run
run:
	@go run main.go

.PHONY: clean
clean:
	@rm -rf vendor
	@rm -rf glide.lock
