package primitive

import "errors"

// FallibleInterface describes values that collect and expose errors.
type FallibleInterface interface {
	ClearErr()
	GetErr() error
	SetErr(err error)
	HasErr() bool
}

// Fallible collects errors produced while processing a value.
type Fallible struct {
	errs []error
}

// SetErr adds a non-nil error to the collection.
func (f *Fallible) SetErr(err error) {
	if err != nil {
		f.errs = append(f.errs, err)
	}
}

// ClearErr removes all collected errors.
func (f *Fallible) ClearErr() {
	f.errs = nil
}

// GetErr returns all collected errors joined together.
func (f *Fallible) GetErr() error {
	return errors.Join(f.errs...)
}

// HasErr reports whether at least one error has been collected.
func (f *Fallible) HasErr() bool {
	return len(f.errs) > 0
}
