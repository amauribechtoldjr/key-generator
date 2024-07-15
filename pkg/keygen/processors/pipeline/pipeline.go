package pipeline

import "github.com/amauribechtoldjr/key-generator/pkg/keygen/processors"

type Pipeline struct {
	pipes []processors.Pipe
}

func New(pipes ...processors.Pipe) processors.Pipe {
	return &Pipeline{pipes: pipes}
}

func (pr Pipeline) Execute(key string) string {
	for _, p := range pr.pipes {
		key = p.Execute(key)
	}

	return key
}