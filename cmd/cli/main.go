package main

import (
	"fmt"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen"
)

func main() {
	flags := keygen.Setup()


	// TODO: this generateKey can be refactored, maybe directly remove it.
	fmt.Println(keygen.GenerateKey(flags)) 
}