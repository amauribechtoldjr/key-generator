package charprocessor

import (
	"math/rand/v2"
	"strings"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors"
)

type CharProcessor struct {
	charCodes []string
	signal    string
}

func New(charCodes ...string) processors.Pipe {
	return &CharProcessor{charCodes: charCodes, signal: "{@}"}
}

func (pc CharProcessor) Execute(key string) string {
	signalsQuantity := strings.Count(key, pc.signal)

	for i := 0; i < signalsQuantity; i++ {
		key = strings.Replace(key, pc.signal, pc.charCodes[rand.IntN(len(pc.charCodes))], 1)
	}

	return key
}