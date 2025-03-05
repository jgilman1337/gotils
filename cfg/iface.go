package cfg

/*
Represents a generic configuration object that has many QoL features like default values,
marshaling/unmarshaling to/from byte streams, and more.
*/
type IConfig[T any] interface {
	// Provides a writable pointer to the config data contained within an IConfig struct.
	Data() *T

	// Sets the default values for this config object.
	Defaults() (IConfig[T], error)

	// Checks if the data of two config objects pointed to by Data() are deeply equal.
	Equal(other IConfig[T]) bool

	// Saves the default config object to the default path.
	// SaveDefault() error
	// Saves a config object to the default path.
	// Save() error
	// Saves this config object to a given path.
	SaveAs(path string) error

	// Saves this config object to the default path.
	// func Load() {}
	// Loads this config object from a given path.
	// func LoadAs() {}

	// Loads this config object from the default path, saving the default version if it doesn't exist.
	// func Init() {}
	// Loads this config object from a given path, saving the default version if it doesn't exist.
	// func InitAs() {}

	// defaultPathName() string //Specifies the default location to drop the default configuration.
}
