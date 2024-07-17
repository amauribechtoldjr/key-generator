package dateprocessor_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/amauribechtoldjr/key-generator/pkg/keygen/processors/dateprocessor"
)

func TestMonths(c *testing.T) {
	isValid := regexp.MustCompile(`test((1[0-2])|(0[1-9]))`)

	pc := dateprocessor.New(time.Now())

	got := pc.Execute("test{mm}")

	if !isValid.MatchString(got) {
		c.Errorf("Months processor failed.")
	}
}

func TestDays(c *testing.T) {
	isValid := regexp.MustCompile(`test((0[1-9]|1[0-9]|2[0-9]|3[0-1]))`)

	pc := dateprocessor.New(time.Now())

	got := pc.Execute("test{dd}")

	if !isValid.MatchString(got) {
		c.Errorf("Days processor failed.")
	}
}

func TestFullYear(c *testing.T) {
	isValid := regexp.MustCompile(`test20([0-9][0-9])`)

	pc := dateprocessor.New(time.Now())

	got := pc.Execute("test{YY}")

	if !isValid.MatchString(got) {
		c.Errorf("Days processor failed.")
	}
}

func TestYear(c *testing.T) {
	isValid := regexp.MustCompile(`test([0-9][0-9])`)

	pc := dateprocessor.New(time.Now())

	got := pc.Execute("test{yy}")

	if !isValid.MatchString(got) {
		c.Errorf("Days processor failed.")
	}
}