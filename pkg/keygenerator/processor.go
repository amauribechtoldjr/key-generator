package keygenerator

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"time"
)

type Processor struct {
	process func(ks *string) *string
	signal string
}

const (
	CHAR_CODE_SIGNAL string = "{@}"
 	YEAR_CODE_SIGNAL string = "{yy}"
 	FULL_YEAR_CODE_SIGNAL string = "{YY}"
)

var fullYear string = strconv.Itoa(time.Now().Year())

func processor(initialKs *string, pcs ...*Processor) string {
	s := initialKs

	for _, pc := range pcs {
		if isValidSignal(s, pc.signal) {
			pc.process(s) 
		}
	}

	return *s
}

func processChar(key *string) *string {
	var CHAR_CODES = [5]string {"!","@","$","%","&"}

	signalsQuantity := strings.Count(*key, CHAR_CODE_SIGNAL)

	for i := 0; i < signalsQuantity; i++ {
		*key = strings.Replace(*key, CHAR_CODE_SIGNAL, CHAR_CODES[rand.IntN(5)], 1)
	}

	return key
}

func processYear(key *string) *string {
	*key = strings.ReplaceAll(*key, YEAR_CODE_SIGNAL, fullYear[0:2])

	return key
}

func processFullYear(key *string) *string {
	*key = strings.ReplaceAll(*key, FULL_YEAR_CODE_SIGNAL, fullYear)

	return key
}

func isValidSignal(key *string, signal string) bool {
	signalIndex := strings.Index(*key, signal)

	return signalIndex > -1
}

func ProcessKeyString(key string) string {
	return processor(
		&key,
		&Processor{ process: processChar, signal: CHAR_CODE_SIGNAL }, 
		&Processor{ process: processYear, signal: YEAR_CODE_SIGNAL },
		&Processor{ process: processFullYear, signal: FULL_YEAR_CODE_SIGNAL },
	)
}