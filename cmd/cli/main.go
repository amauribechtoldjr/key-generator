package main

import (
	"fmt"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen"
)

func main() {
	flags := Setup()

	fmt.Println(keygen.ProcessKeyString(flags.Ks))
}