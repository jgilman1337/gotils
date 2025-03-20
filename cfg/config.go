package cfg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"

	"github.com/creasty/defaults"
	"github.com/jgilman1337/gotils/cfg/marshaler"
)

var (
	ErrNoMarshalers = errors.New("cannot marshal/unmarshal; no marshalers are bound to this object")
	ErrFloadFailure = errors.New("failed to load the config file pointed to by the current marshaller")
	ErrMismatchedBM = errors.New("mismatched marshaler and in byte array sizes")
	ErrReadFailure  = errors.New("failed to read the config bytes provided by the current file")
)

// Defines a default configuration provider function.
type DefaultsProvider[T any] func() (*T, error)

// Implements a basic configuration object that contains a data struct, which holds thr actual configuration data.
type Config[T any] struct {
	DFunc      DefaultsProvider[T]   //The function that will set default values.
	data       T                     //The inner configuration data object.
	marshalers []marshaler.Marshaler //The config marshalers that are bound to this object.
}

//TODO: try to use a custom struct tag `kname` that indicates the name of the key

// Enforces compliance with the IConfig interface.
var _ IConfig[any] = (*Config[any])(nil)

// Creates a new Config object using a data struct.
func NewConfig[T any](data T) *Config[T] {
	return &Config[T]{
		data:       data,
		marshalers: make([]marshaler.Marshaler, 0),
	}
}

/*
Creates a new Config object with the default values of the bound struct in the type.
This function equivalent to the following statement:

	cfg := must(NewConfig(cfgObject{}).Defaults())
*/
func NewConfigDefaults[T any]() *Config[T] {
	var dat T
	out, err := NewConfig(dat).Defaults()
	if err != nil {
		panic(err)
	}
	return out.(*Config[T])
}

// Implements the BindMarshaler() function from IConfig.
func (c *Config[T]) BindMarshaler(nms ...marshaler.Marshaler) IConfig[T] {
	c.marshalers = mergeMarshalerArrays(c.marshalers, nms)
	return c
}

// Implements the Data() function from IConfig.
func (c *Config[T]) Data() *T {
	return &c.data
}

// Implements the Defaults() function from IConfig. Uses creasty/defaults or a custom provider to provide the default object.
func (c *Config[T]) Defaults() (IConfig[T], error) {
	//Use the user provided defaults provider if set
	if c.DFunc != nil {
		//Get the default object from the user provided function
		defaults, err := c.DFunc()
		if err != nil {
			return nil, err
		}
		c.data = *defaults
		return c, nil
	} else {
		return c, defaults.Set(&c.data)
	}
}

// Implements the Equal() function from IConfig.
func (c Config[T]) Equal(other IConfig[T]) bool {
	return reflect.DeepEqual(c.Data(), other.Data())
}

// Implements the LoadBytes() function from IConfig.
func (c *Config[T]) LoadBytes(in ...[]byte) (*T, error) {
	//Ensure the number of marshalers and varargs match
	if len(in) != len(c.marshalers) {
		return nil, fmt.Errorf("%w; marshalers: %d, byte arrays: %d", ErrMismatchedBM, len(c.marshalers), len(in))
	}

	//Unmarshal the byte arrays in order of appearance
	for i, m := range c.marshalers {
		if err := m.UMarshal(in[i], &(c.data)); err != nil {
			return nil, fmt.Errorf("config (load bytes): feeder error: %v", err)
		}
	}

	return &c.data, nil
}

// Implements the LoadPath() function from IConfig.
func (c *Config[T]) LoadPath() (*T, error) {
	//Loop over the available marshalers
	for _, m := range c.marshalers {
		//Read in the file pointed to by the marshaler
		mf, err := os.Open(m.Path())
		if err != nil {
			return nil, fmt.Errorf("%w; marshaler: %T, file: %s", ErrFloadFailure, m, err)
		}
		mb, err := io.ReadAll(mf)
		if err != nil {
			return nil, fmt.Errorf("%w; file: %s", ErrReadFailure, m.Path())
		}

		//Unmarshal the bytes to a struct
		if err := m.UMarshal(mb, &(c.data)); err != nil {
			return nil, fmt.Errorf("config (load file): feeder error: %v", err)
		}
	}

	return &c.data, nil
}

// Implements the Save() function from IConfig.
func (c Config[T]) Save() error {
	//Ensure at least one marshaler is bound before continuing
	if len(c.marshalers) == 0 {
		return ErrNoMarshalers
	}

	//Run each marshaler
	for _, mar := range c.marshalers {
		//Run the current marshaler, but only if it has a set path value
		if mar.Path() != "" {
			data, err := mar.Marshal(&c.data)
			if err != nil {
				return fmt.Errorf("config (marshal %s): failed to marshal; %w", reflect.TypeOf(mar).Name(), err)
			}

			err = os.WriteFile(filepath.Clean(mar.Path()), data, 0644)
			if err != nil {
				return fmt.Errorf("config (write %s) failed to write to file %s; %w", reflect.TypeOf(mar).Name(), mar.Path(), err)
			}
		}
	}

	return nil
}
