package marshaler

import (
	"encoding/json"
)

// Represents a JSON marshaler that implements Marshaler.
type Json[T any] struct {
	priority int8
}

// Enforces compliance with the Marshaler interface.
var _ Marshaler[any] = (*Json[any])(nil)

// Creates a new Json marshaler object with a the default priority of 0.
func NewJson[T any]() Json[T] {
	return Json[T]{priority: 0}
}

// Creates a new Json marshaler object with a given priority.
func NewJsonPriority[T any](priority int8) Json[T] {
	return Json[T]{priority: priority}
}

// Implements the BackedByFile() function from Marshaler.
func (j Json[T]) BackedByFile() bool {
	return true
}

// Implements the DefaultPath() function from Marshaler.
func (j Json[T]) DefaultPath() string {
	return "config.json"
}

// Implements the Ident() function from Marshaler.
func (j Json[T]) Ident() string {
	return "json_marshaler"
}

// Implements the Marshal() function from Marshaler.
func (j Json[T]) Marshal(c *T) ([]byte, error) {
	return json.Marshal(c)
}

// Implements the Priority() function from Marshaler.
func (j Json[T]) Priority() int8 {
	return j.priority
}

// Implements the UMarshal() function from Marshaler.
func (j Json[T]) UMarshal(b []byte, c *T) error {
	return json.Unmarshal(b, &c)
}
