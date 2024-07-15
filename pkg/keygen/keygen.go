package keygen

import (
	"time"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors/charprocessor"
	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors/dateprocessor"
	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors/pipeline"
)


func ProcessKeyString(key string) string {
	return pipeline.New(
		charprocessor.New("!", "@", "$", "%", "&", "#"),
		dateprocessor.New(time.Now()),
		).Execute(key)
	}

const (
	FULL_YEAR_CODE_SIGNAL     string = "{YY}"
	DAYS_CODE_SIGNAL          string = "{dd}"
	MONTHS_CODE_SIGNAL        string = "{mm}"
	RANGE_NUMBERS_CODE_SIGNAL string = "{n:9}"
)


// func processFullYear(key *string) *string {
// 	// fmt.Println(&key)
// 	
// }

// func processMonths(key *string) *string {
// 	// fmt.Println(&key)
// 	*key = strings.ReplaceAll(*key, MONTHS_CODE_SIGNAL, time.Now().Format("01"))

// 	return key
// }

// func processDays(key *string) *string {
// 	// fmt.Println(&key)
// 	*key = strings.ReplaceAll(*key, DAYS_CODE_SIGNAL, time.Now().Format("02"))

// 	return key
// }

// func isValidSignal(key *string, signal string) bool {
// 	// fmt.Println(&key)
// 	signalIndex := strings.Index(*key, signal)

// 	return signalIndex > -1
// }