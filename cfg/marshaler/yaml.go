package marshaler

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// Represents a YAML marshaler that implements Marshaler.
type Yaml struct {
	path string
}

// Enforces compliance with the Marshaler interface.
var _ Marshaler = (*Yaml)(nil)

// Creates a new Yaml marshaler object.
func NewYaml(path string) Yaml {
	if path == "" {
		path = "config.yml"
	}
	return Yaml{path: path}
}

// Implements the Marshal() function from Marshaler.
func (y Yaml) Marshal(c any) ([]byte, error) {
	var out bytes.Buffer
	if err := yaml.NewEncoder(&out).Encode(c); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// Implements the Path() function from Marshaler.
func (y Yaml) Path() string {
	return y.path
}

// Implements the UMarshal() function from Marshaler.
func (y Yaml) UMarshal(b []byte, c any) error {
	return yaml.NewDecoder(bytes.NewBuffer(b)).Decode(c)
}
