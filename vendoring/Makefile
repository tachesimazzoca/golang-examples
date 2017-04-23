GO_EXECUTABLE ?= go
GOPATH := $(CURDIR)

define fetch_git_repo
  if [ ! -e $2 ]; \
  then \
    git clone $1 $2; \
  else \
    cd $2 && git checkout master && git pull; \
  fi
endef

.PHONY: install-deps
install-deps:
	@$(call fetch_git_repo, \
  "https://github.com/mitchellh/colorstring.git", \
  "src/vendor/github.com/mitchellh/colorstring")

.PHONY: run
run:
	@$(GO_EXECUTABLE) run src/vendoring/main.go

.PHONY: install
install:
	@$(GO_EXECUTABLE) install vendoring

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf src/vendor/github.com
