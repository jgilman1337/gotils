package marshaler

import "github.com/jgilman1337/gotils/cfg"

// Defines the structure of a function that marshals a config object to a byte stream.
type BMarshalerFunc func(c *cfg.IConfig) ([]byte, error)

// Defines the structure of a function that unmarshals a config object from a byte stream.
type BUMarshalerFunc func(b []byte, c *cfg.IConfig) error
