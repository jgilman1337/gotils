package marshaler

import "github.com/jgilman1337/gotils/cfg"

// Defines the structure of a function that marshals a config object to a byte stream.
type MarshalerFunc func(c cfg.IConfig) ([]byte, error)

// Defines the structure of a function that unmarshals a config object from a byte stream.
type UMarshalerFunc func(b []byte, c cfg.IConfig) error

// Represents a basic marshal and unmarshal object that can be used to read/write to/from config files.
type Marshaler interface {
	// The default path	to use when reading/writing.
	DefaultPath() string
	
	// Converts a config object to a byte stream for writing to a file.
	MFunc(c cfg.IConfig) ([]byte, error)
	
	// Indicates the priority of the marshaler. Higher numbers generally mean higher priority and run last when marshaling and unmarshaling.
	Priority() int
	
	// Converts a byte stream to a config object for usage.
	UFunc(b []byte, c cfg.IConfig) error
}