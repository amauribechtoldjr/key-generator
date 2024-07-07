package generator

import (
	"math/rand/v2"
	"strings"
)

const CHAR_CODE_SIGNAL string = "{@}"

func processChar(key *string) {
	specialCharIndex := strings.Index(*key, CHAR_CODE_SIGNAL)

	if specialCharIndex == -1 {
		return
	}

	var CHAR_CODES = [5]string {"!","@","$","%","&"}

	*key = strings.ReplaceAll(*key, CHAR_CODE_SIGNAL, CHAR_CODES[rand.IntN(5)])
}

func ProcessKeyString(key string) string {
	myKeyString := key

	processChar(&myKeyString)

	return myKeyString
}