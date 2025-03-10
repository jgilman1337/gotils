package marshaler

import (
	"encoding/json"

	"github.com/jgilman1337/gotils/cfg/iface"
)

// Represents a JSON marshaler that implements Marshaler.
type Json[T any] struct{}

// Enforces compliance with the Marshaler interface.
var _ iface.Marshaler[any] = (*Json[any])(nil)

// Implements the DefaultPath() function from Marshaler.
func (j Json[T]) DefaultPath() string {
	return "config.json"
}

// Implements the Marshal() function from Marshaler.
func (j Json[T]) Marshal(c iface.IConfig[T]) ([]byte, error) {
	return json.Marshal(c.Data())
}

// Implements the Priority() function from Marshaler.
func (j Json[T]) Priority() int {
	return 0
}

// Implements the UMarshal() function from Marshaler.
func (j Json[T]) UMarshal(b []byte, c iface.IConfig[T]) error {
	return json.Unmarshal(b, c.Data())
}
