package marshaler

import (
	"encoding/json"
)

// Represents a JSON marshaler that implements Marshaler.
type Json struct {
	path     string
	Minified bool
}

// Enforces compliance with the Marshaler interface.
var _ Marshaler = (*Json)(nil)

// Creates a new Json marshaler object with a the default priority of 0.
func NewJson(path string) Json {
	if path == "" {
		path = "config.json"
	}
	return Json{path: path}
}

// Implements the Marshal() function from Marshaler.
func (j Json) Marshal(c any) ([]byte, error) {
	if j.Minified {
		return json.Marshal(c)
	}
	return json.MarshalIndent(c, "", "\t")
}

// Implements the Path() function from Marshaler.
func (j Json) Path() string {
	return j.path
}

// Implements the UMarshal() function from Marshaler.
func (j Json) UMarshal(b []byte, c any) error {
	return json.Unmarshal(b, c)
}
