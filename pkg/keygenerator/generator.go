package keygenerator

import (
	"fmt"
)

func GenerateKey() {
	flags := Setup()

	fmt.Println(ProcessKeyString(flags.Ks)) 
}