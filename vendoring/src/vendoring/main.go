package main

import (
	"hello"
)

func main() {
	hello.Success("%s\n", "Hello, vendoring!")
	hello.Failure("%s\n", "Does it mean vending or vendorizing?")
}
