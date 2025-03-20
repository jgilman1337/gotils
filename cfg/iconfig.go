package cfg

import "github.com/jgilman1337/gotils/cfg/marshaler"

/*
Represents a generic configuration object that has many QoL features like default values,
marshaling/unmarshaling to/from byte streams, and more.
*/
type IConfig[T any] interface {
	// Binds a marshaler or list of marshalers to this config object. Marshalers loaded last gain higher priority.
	BindMarshaler(nms ...marshaler.Marshaler)

	// Provides a writable pointer to the config data contained within an IConfig struct.
	Data() *T

	// Sets the default values for this config object.
	Defaults() (IConfig[T], error)

	// Checks if the data of two config objects pointed to by Data() are deeply equal.
	Equal(other IConfig[T]) bool

	// Loads this config object from the path(s) specified by the bound marshalers, saving the default version if it doesn't exist.
	// Init() {}

	// Loads this config object from given input bytes. The length of the varargs must match the number of bound marshalers.
	LoadBytes(in ...[]byte) error

	// Loads this config object from the path(s) specified by the bound marshalers.
	LoadPath() error

	// Saves the default config object to the default path.
	// SaveDefault() error
	// Saves a config object to the path(s) specified by the bound marshalers.
	Save() error
}
