package cfg

import (
	//"encoding/json"
	"fmt"
	"github.com/creasty/defaults"
	//"os"
	"reflect"
)

// Defines a default configuration provider function.
type DefaultsProvider func() (IConfig, error)

// Impelments a basic configuration object that contains a data struct, which holds thr actual configuration data.
type Config[T any] struct {
	data T // The inner configuration data object.

	DFunc DefaultsProvider // The function that will set default values.
}

// Enforces compliance with the IConfig interface.
var _ IConfig = (*Config[any])(nil)

// Creates a new Config object using a data struct.
func NewConfig[T any](data T) *Config[T] {
	return &Config[T]{
		data: data,
	}
}

// Implements the Data() function from IConfig.
func (c *Config[T]) Data() any {
	return c.data //TODO: maybe use generics here to avoid unnecessary casts
}

// Implements the Defaults() function from IConfig. Uses creasty/defaults or a custom provider to provide the default object.
func (c *Config[T]) Defaults() (IConfig, error) {
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
