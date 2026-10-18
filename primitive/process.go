package primitive

import interfass "github.com/cyril-fl/plum/interface"

// Interface

type Processable interface {
	FallibleInterface
	interfass.Runnable
}

type ProcessableResult[T any] interface {
	Processable

	GetResult() T
}

type Process struct {
	Fallible
}

// Methods
func NewErr() *Process {
	return &Process{}
}

// Helpers
func HandleProcess(processes []Processable) error {
	for _, p := range processes {
		p.Run()

		if p.HasErr() {
			return p.GetErr()
		}
	}

	return nil
}

/*
NOTE Snippet - Process

type p struct {
	process.Process
	cache
}

type cache struct {
}

func Process() *p {
	return &p{

	}
}

func (i *p) Run() {}

// Getters & Setters
func (i *p) GetResult() {
	logs.SilentWarning("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "",
	})
}
// Methods

// Helpers

*/
