# golang-examples/vendoring

## Overview

This example shows how to resolve packages under the `vendor` directory in Golang.

```
# Install dependencies into vendor
$ make install-deps

# Run src/vendoring/main.go
$ make run
...

# Install binaries
$ make install
$ bin/vendoring
...

# Remove all builds and cloned depedencies
$ make clean
```

* The `GOPATH` must be defined for this _vendoring_ feature.
* Any programs must be put under a package directory in `$GOPATH/src`, such as `src/foo/main.go`.

