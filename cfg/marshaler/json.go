package marshaler

import (
	"encoding/json"

	"github.com/jgilman1337/gotils/cfg"
)

// Represents a JSON marshaler that implements Marshaler.
type Json[T any] struct{}

// Enforces compliance with the Marshaler interface.
var _ Marshaler[any] = (*Json[any])(nil)

// Implements the DefaultPath() function from Marshaler.
func (j Json[T]) DefaultPath() string {
	return "config.json"
}

// Implements the MFunc() function from Marshaler.
func (j Json[T]) MFunc(c cfg.IConfig[T]) ([]byte, error) {
	return json.Marshal(c.Data())
}

// Implements the Priority() function from Marshaler.
func (j Json[T]) Priority() int {
	return 0
}

// Implements the UFunc() function from Marshaler.
func (j Json[T]) UFunc(b []byte, c cfg.IConfig[T]) error {
	return json.Unmarshal(b, c.Data())
}
