package cfg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/creasty/defaults"
	"github.com/jgilman1337/gotils/cfg/marshaler"
)

var (
	ErrNoMarshalers          = errors.New("cannot marshal/unmarshal; no marshalers are bound to this object")
	ErrMarshalerAlreadyBound = errors.New("the incoming marshaler at pos %d is already bound as %s (priority: %d)")
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

// Creates a new Config object with the default values of the bound struct in the type.
func NewConfigDefaults[T any]() *Config[T] {
	var dat T
	out, err := NewConfig(dat).Defaults()
	if err != nil {
		panic(err)
	}
	return out.(*Config[T])
}

// Implements the BindMarshaler() function from IConfig.
func (c *Config[T]) BindMarshaler(nms ...marshaler.Marshaler) {
	c.marshalers = mergeMarshalerArrays(c.marshalers, nms)
}

// Implements the Data() function from IConfig.
func (c *Config[T]) Data() *T {
	return &c.data
}

// Implements the Defaults() function from IConfig. Uses creasty/defaults or a custom provider to provide the default object.
func (c *Config[T]) Defaults() (IConfig[T], error) {
	//Use the defaults provider if set
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
				return fmt.Errorf("(%s) failed to marshal; %w", reflect.TypeOf(mar).Name(), err)
			}

			err = os.WriteFile(filepath.Clean(mar.Path()), data, 0644)
			if err != nil {
				return fmt.Errorf("(%s) failed to write to file %s; %w", reflect.TypeOf(mar).Name(), mar.Path(), err)
			}
		}
	}

	return nil
}
