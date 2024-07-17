package charprocessor_test

import (
	"regexp"
	"testing"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors/charprocessor"
)

func TestSingleChar(c *testing.T) {
	isValid := regexp.MustCompile(`test(@|!|\$|%|&|#)`)

	pc := charprocessor.New("!", "@", "$", "%", "&", "#")

	got := pc.Execute("test{@}")

	if !isValid.MatchString(got) {
    c.Errorf("Char processor test failed!")
	}
}

func TestMultipleChar(c *testing.T) {
	isValid := regexp.MustCompile(`test(@|!|\$|%|&|#)test(@|!|\$|%|&|#)`)

	pc := charprocessor.New("!", "@", "$", "%", "&", "#")

	got := pc.Execute("test{@}test{@}")

	if !isValid.MatchString(got) {
    c.Errorf("Char processor test failed!")
	}
}
