package primitive

import interfass "github.com/cyril-fl/plum/interface"

// Processable describes a runnable value that can report errors.
type Processable interface {
	FallibleInterface
	interfass.Runnable
}

// ProcessableResult describes a processable value that exposes a result.
type ProcessableResult[T any] interface {
	Processable

	GetResult() T
}

// Process stores errors produced by a process.
type Process struct {
	Fallible
}

// NewErr returns an empty Process ready to collect errors.
func NewErr() *Process {
	return &Process{}
}

// HandleProcess runs processes in order and returns the first error encountered.
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

func (i *p) GetResult() {
	logs.SilentWarning("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "",
	})
}
// Helpers

*/
