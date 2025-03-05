package cfg

import (
	//"encoding/json"
	"fmt"

	"github.com/creasty/defaults"

	//"os"
	"reflect"
)

// Defines a default configuration provider function.
type DefaultsProvider[T any] func() (IConfig[T], error)

// Implements a basic configuration object that contains a data struct, which holds thr actual configuration data.
type Config[T any] struct {
	data  T                   //The inner configuration data object.
	DFunc DefaultsProvider[T] //The function that will set default values.
}

//TODO: try to use a custom struct tag `kname` that indicates the name of the key

// Enforces compliance with the IConfig interface.
var _ IConfig[any] = (*Config[any])(nil)

// Creates a new Config object using a data struct.
func NewConfig[T any](data T) *Config[T] {
	return &Config[T]{
		data: data,
	}
}

// Implements the Data() function from IConfig.
func (c *Config[T]) Data() *T {
	return &c.data
}

// Implements the Defaults() function from IConfig. Uses creasty/defaults or a custom provider to provide the default object.
func (c *Config[T]) Defaults() (IConfig[T], error) {
	//return &Config[T]{Data: Zero[T]()}, nil

	//Use the defaults provider if set
	if c.DFunc != nil {
		//Get the default values
		defaults, err := c.DFunc()
		if err != nil {
			return nil, err
		}

		//Use reflection to set the value
		//TODO: make this a utility
		v := reflect.ValueOf(c)
		if v.Kind() != reflect.Ptr || v.IsNil() {
			return nil, fmt.Errorf("c must be a non-nil pointer")
		}

		//Ensure the types match
		dv := reflect.ValueOf(defaults)
		if v.Elem().Type() != dv.Type() {
			return nil, fmt.Errorf("type mismatch: cannot assign %v to %v", dv.Type(), v.Elem().Type())
		}

		//Set defaults
		v.Elem().Set(dv)
		return c, nil
	} else {
		return c, defaults.Set(&c.data)
	}
}

// Implements the Equal() function from IConfig.
func (c Config[T]) Equal(other IConfig[T]) bool {
	return reflect.DeepEqual(c.Data(), other.Data())
}

// Implements the SaveAs() function from IConfig.
func (c Config[T]) SaveAs(path string) error {
	/*
		data, err := json.Marshal(c.Data)
		if err != nil {
			return err
		}
		return os.WriteFile(path, data, 0644)
	*/
	return nil
}
