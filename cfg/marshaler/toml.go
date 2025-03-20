package marshaler

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// Represents a TOML marshaler that implements Marshaler.
type Toml struct {
	path string
}

// Enforces compliance with the Marshaler interface.
var _ Marshaler = (*Toml)(nil)

// Creates a new Toml marshaler object.
func NewToml(path string) Toml {
	if path == "" {
		path = "config.toml"
	}
	return Toml{path: path}
}

// Implements the Marshal() function from Marshaler.
func (t Toml) Marshal(c any) ([]byte, error) {
	var out bytes.Buffer
	enc := toml.NewEncoder(&out)
	enc.Indent = "\t"
	if err := enc.Encode(c); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// Implements the Path() function from Marshaler.
func (t Toml) Path() string {
	return t.path
}

// Implements the UMarshal() function from Marshaler.
func (t Toml) UMarshal(b []byte, c any) error {
	_, err := toml.NewDecoder(bytes.NewBuffer(b)).Decode(c)
	return err
}
