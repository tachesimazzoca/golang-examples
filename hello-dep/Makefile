.PHONY: install-deps
install-deps:
	@dep ensure

.PHONY: run
run:
	@go run main.go

.PHONY: clean
clean:
	@rm -rf vendor
	@rm -rf Gopkg.lock
