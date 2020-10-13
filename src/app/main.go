package main

import (
	"fmt"

	"github.com/alitari/kogo/src/app/mod1"
	sample "github.com/elliotforbes/test-package"
)

func main() {
	fmt.Printf("Hello, im app\n")
	mod1.Hello()
	sample.MySampleFunc()
}
