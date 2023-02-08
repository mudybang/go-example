package main

import (
	"fmt"

	"github.com/mudybang/go-example/hello"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(hello.Go())
}
