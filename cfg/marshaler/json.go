package marshaler

import (
	"encoding/json"
	"github.com/jgilman1337/gotils/cfg"
)

// Represents a JSON marshaler that implements Marshaler.
type Json struct {}

// Enforces compliance with the Marshaler interface.
var _ Marshaler = (*Json)(nil)

// Implements the DefaultPath() function from Marshaler.
func (j Json) DefaultPath() string {
	return "config.json"
}

// Implements the MFunc() function from Marshaler.
func (j Json) MFunc(c cfg.IConfig) ([]byte, error) {
	return json.Marshal(c.Data()) //TODO: data pass here
}

// Implements the Priority() function from Marshaler.
func (j Json) Priority() int {
	return 0
}

// Implements the UFunc() function from Marshaler.
func (j Json) UFunc(b []byte, c cfg.IConfig) error {
	return json.Unmarshal(b, &c.Data())
}