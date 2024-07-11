package main

import (
	"fmt"

	"github.com/amauribechtoldjr/key-generator/pkg/keygenerator"
)

func main() {
	flags := keygenerator.Setup()

	fmt.Println(keygenerator.GenerateKey(flags)) 
}