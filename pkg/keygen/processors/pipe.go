package processors

type Pipe interface {
	Execute(string) string
}