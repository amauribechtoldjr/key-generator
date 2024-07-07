package generator

import (
	"errors"
	"flag"
)

type KeyGeneratorFlags struct {
	Ks string
}

func ReadFlags() (KeyGeneratorFlags, error) {
	keyStringFlag := flag.String("keyString", "", "Your key string generator")
	ksFlag := flag.String("ks", "", "Your key string generator")

	flag.Parse()

	if *ksFlag != "" {
		return KeyGeneratorFlags{Ks: *ksFlag}, nil
	}

	if *keyStringFlag != "" {
		return KeyGeneratorFlags{Ks: *keyStringFlag}, nil
	}

	return KeyGeneratorFlags{}, errors.New("you should provide your key string")
}

