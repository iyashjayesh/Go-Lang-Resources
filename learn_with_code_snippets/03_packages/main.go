package main

import (
	"fmt"
	"math" // this is the package

	"github.com/iyashjayesh/go-lang-resources/learn_with_code_snippets/03_packages/strutil" // this is the package
)

func main() {
	fmt.Println(math.Floor(22.12))
	fmt.Println(math.Ceil(22.12))
	fmt.Println(math.Sqrt(9))

	fmt.Println("Calling the function from another package")
	fmt.Println(strutil.Reverse("Hello"))

}
