package keygenerator

import (
	"errors"
	"flag"
)

type KeyGeneratorFlags struct {
	Ks string
}

type SetOptionFn func(*KeyGeneratorFlags) error

func New(ff ...SetOptionFn) (*KeyGeneratorFlags, error) {
	o := &KeyGeneratorFlags{}

	for _, f := range ff {
		if err := f(o); err != nil {
			return nil, err
		}
	}

	return o, nil
}

func Must(ff ...SetOptionFn) *KeyGeneratorFlags {
	o, err := New(ff...)

	if err!= nil {
		panic(err)
	}
	
	return o
}

func WithFlag(key, usage string) SetOptionFn {
	return func(o *KeyGeneratorFlags) error {
		flag.StringVar(&o.Ks, key, "", usage)

		return nil
	}
}

func InitFlags() SetOptionFn {
	return func(_ *KeyGeneratorFlags) error {
		flag.Parse()

		return nil
	}
}

func WithValidKS() SetOptionFn {
	return func(o *KeyGeneratorFlags) error {
		if o.Ks == "" {
			return errors.New("should provide a key string")
		}
		
		return nil
	}
}

func Setup() *KeyGeneratorFlags {
	const usage = "Your key string generator"

	return Must(
		WithFlag("key-string", usage),
		WithFlag("ks", usage + " (shorthand)"),
		InitFlags(),
		WithValidKS(),
	)
}

 