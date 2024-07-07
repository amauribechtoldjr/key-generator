package generator

import "fmt"

func GenerateKey() {
	if flags, errors := ReadFlags(); errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(ProcessKeyString(flags.Ks)) 
	}
}