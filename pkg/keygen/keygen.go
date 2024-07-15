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