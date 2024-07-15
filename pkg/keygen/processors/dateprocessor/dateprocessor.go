package dateprocessor

import (
	"strconv"
	"strings"
	"time"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors"
)

type DateProcessor struct {
	yearSignal    	string
	fullYearSignal 	string
	daysSignal 			string
	monthSignal 		string
	date 						time.Time
}

func New(date time.Time) processors.Pipe {
	return &DateProcessor{
		yearSignal: "{yy}", 
		fullYearSignal: "{YY}", 
		daysSignal: "{dd}", 
		monthSignal: "{mm}", 
		date: date,
	}
}

func (dp DateProcessor) Execute(key string) string {
	key = strings.ReplaceAll(key, dp.yearSignal, strconv.Itoa(dp.date.Year())[0:2])
	key = strings.ReplaceAll(key, dp.fullYearSignal, strconv.Itoa(dp.date.Year()))
	key = strings.ReplaceAll(key, dp.monthSignal, dp.date.Format("01"))
	key = strings.ReplaceAll(key, dp.daysSignal, dp.date.Format("02"))

	return key
}