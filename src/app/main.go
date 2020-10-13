package main

import (
	"app/mod1"
	"fmt"

	sample "github.com/elliotforbes/test-package"
)

func main() {
	fmt.Printf("Hello, im app\n")
	mod1.Hello()
	sample.MySampleFunc()
}
