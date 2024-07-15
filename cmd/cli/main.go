package main

import (
	"fmt"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen"
)

func main() {
	flags := keygen.Setup()

	fmt.Println(keygen.GenerateKey(flags)) 
}