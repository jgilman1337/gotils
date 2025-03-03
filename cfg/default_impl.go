package cfg

import (
	"encoding/json"
	"os"
)

// Base configuration struct with generic type parameter
type Config[T any] struct {
	Data T
}

// Verify interface implementation
var _ IConfig = (*Config[any])(nil)

func NewConfig[T any](data T) *Config[T] {
	return &Config[T]{
		Data: data,
	}
}

// Parent-implemented Save method
func (c *Config[T]) SaveAs(path string) error {
	data, err := json.Marshal(c.Data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Parent-implemented Defaults method (works for any type)
func (c *Config[T]) Defaults() (IConfig, error) {
	//return &Config[T]{Data: Zero[T]()}, nil
	return DefaultsHelper(c, nil)
}
