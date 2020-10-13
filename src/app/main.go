package main

import (
	"fmt"
	"mod1"

	sample "github.com/elliotforbes/test-package"
)

func main() {
	fmt.Printf("Hello, im app\n")
	mod1.Hello()
	sample.MySampleFunc()
}
