package cfg

import (
	"fmt"
	"reflect"

	"github.com/creasty/defaults"
)

// Defines a default configuration provider function. As long as the returned type correctly implements `IConfig`, any struct may be used.
type DefaultsProvider func() (IConfig, error)

// Helper function for configuration types implementing `IConfig` that uses `creasty/defaults` or a custom provider to provide a default object.
func DefaultsHelper(c IConfig, dfunc DefaultsProvider) (IConfig, error) {
	//Use the defaults provider if set
	if dfunc != nil {
		//Get the default values
		defaults, err := dfunc()
		if err != nil {
			return nil, err
		}

		//Use reflection to set the value
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
		return c, defaults.Set(c)
	}
}

// Helper function to get zero value of any type
func Zero[T any]() T {
	var zero T
	return zero
}
