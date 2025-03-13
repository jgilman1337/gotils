package marshaler

// Represents a basic marshal and unmarshal object that can be used to read/write to/from config files.
type Marshaler interface {
	// Converts a config object to a byte stream for writing to a file.
	Marshal(c any) ([]byte, error)

	// The path	to use when reading/writing.
	Path() string

	// Converts a byte stream to a config object for usage.
	UMarshal(b []byte, c any) error
}
