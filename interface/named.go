// Package interfass defines small behavioral interfaces.
package interfass

// Named is implemented by values that expose a name.
type Named interface {
	GetName() string
}
