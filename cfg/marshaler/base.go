package marshaler

import "github.com/jgilman1337/gotils/cfg"

// Defines the structure of a function that marshals a config object to a byte stream.
type MarshalerFunc[T any] func(c cfg.IConfig[T]) ([]byte, error)

// Defines the structure of a function that unmarshals a config object from a byte stream.
type UMarshalerFunc[T any] func(b []byte, c *cfg.IConfig[T]) error

// Represents a basic marshal and unmarshal object that can be used to read/write to/from config files.
type Marshaler[T any] interface {
	// The default path	to use when reading/writing.
	DefaultPath() string

	// Converts a config object to a byte stream for writing to a file.
	MFunc(c cfg.IConfig[T]) ([]byte, error)

	// Indicates the priority of the marshaler. Higher numbers generally mean higher priority and run last when marshaling and unmarshaling.
	Priority() int

	// Converts a byte stream to a config object for usage.
	UFunc(b []byte, c cfg.IConfig[T]) error
}
