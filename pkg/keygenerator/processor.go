package keygenerator

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"time"
)

const (
	CHAR_CODE_SIGNAL string = "{@}"
 	YEAR_CODE_SIGNAL string = "{yy}"
 	FULL_YEAR_CODE_SIGNAL string = "{YY}"
)

func processChar(key *string) {
	if !isValidSignal(key, CHAR_CODE_SIGNAL) {
		return 
	}

	var CHAR_CODES = [5]string {"!","@","$","%","&"}

	signalsQuantity := strings.Count(*key, CHAR_CODE_SIGNAL)

	for i := 0; i < signalsQuantity; i++ {
		*key = strings.Replace(*key, CHAR_CODE_SIGNAL, CHAR_CODES[rand.IntN(5)], 1)
	}
}

func processYear(key *string) {
	if !isValidSignal(key, FULL_YEAR_CODE_SIGNAL) && !isValidSignal(key, YEAR_CODE_SIGNAL) {
		return 
	}

	fullYear := strconv.Itoa(time.Now().Year())

	*key = strings.ReplaceAll(*key, FULL_YEAR_CODE_SIGNAL, fullYear)
	*key = strings.ReplaceAll(*key, YEAR_CODE_SIGNAL, fullYear[0:2])
}

func isValidSignal(key *string, signal string) bool {
	signalIndex := strings.Index(*key, signal)

	return signalIndex > -1
}

func ProcessKeyString(key string) string {
	myKeyString := key

	processChar(&myKeyString)
	processYear(&myKeyString)

	return myKeyString
}