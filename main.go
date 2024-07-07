package main

import (
	"fmt"

	"github.com/amauribechtoldjr/key-generator/generator"
)

func main() {
	if flags, errors := generator.ReadFlags(); errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(flags.Ks)
	}
}