package cfg

import "github.com/jgilman1337/gotils/cfg/marshaler"

// Creates a new config object using a JSON file as its backing marshaler.
func NewWithJson[T any]() *Config[T] {
	jm := marshaler.NewJson[T]()
	out := presetCreate[T]()
	out.marshalers.Enqueue(jm)
	return out
}

// Creates new config instances for the preset providers.
func presetCreate[T any]() *Config[T] {
	//Get the zero value of the generic datatype
	var zero T

	//Initialize the config object
	cfg, err := NewConfig(zero).Defaults()
	if err != nil {
		panic(err)
	}

	return cfg.(*Config[T])
}
