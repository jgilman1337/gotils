package cfg

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/creasty/defaults"
	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
	"github.com/jgilman1337/gotils/cfg/marshaler"
)

var (
	ErrNoMarshalers          = errors.New("cannot marshal/unmarshal; no marshalers are bound to this object")
	ErrMarshalerAlreadyBound = errors.New("the incoming marshaler at pos %d is already bound as %s (priority: %d)")
)

// Defines a default configuration provider function.
type DefaultsProvider[T any] func() (IConfig[T], error)

// Implements a basic configuration object that contains a data struct, which holds thr actual configuration data.
type Config[T any] struct {
	DFunc      DefaultsProvider[T]               //The function that will set default values.
	data       T                                 //The inner configuration data object.
	marshalers *pq.Queue[marshaler.Marshaler[T]] //The config marshalers that are bound to this object.
	//TODO: use a custom struct or lookup table so it can be determined whether the marshaler has a backing file and more; don't just use a priority queue here
}

//TODO: try to use a custom struct tag `kname` that indicates the name of the key

// Enforces compliance with the IConfig interface.
var _ IConfig[any] = (*Config[any])(nil)

// Creates a new Config object using a data struct.
func NewConfig[T any](data T) *Config[T] {
	//Setup the comparator function for the marshalers
	mcomparator := func(a, b marshaler.Marshaler[T]) int {
		priA, priB := a.Priority(), b.Priority()
		switch {
		case priA > priB:
			return 1
		case priA < priB:
			return -1
		default:
			return 0
		}
	}

	return &Config[T]{
		data:       data,
		marshalers: pq.NewWith(mcomparator),
	}
}

// Implements the BindMarshaler() function from IConfig.
func (c *Config[T]) BindMarshaler(nms ...marshaler.Marshaler[T]) error {
	for i, nm := range nms {
		for _, em := range c.marshalers.Values() {
			//Check if one of the incoming marshalers is already bound
			cmid := em.Ident()
			if nm.Ident() == cmid {
				return fmt.Errorf(strconv.Itoa(i), cmid, em.Priority())
			}

			//No clashes; add the marshaler
			c.marshalers.Enqueue(nm)
		}
	}

	return nil
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
		//TODO: may not be necessary to do all of these checks; more testing is required
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

// Implements the LoadAs() function from IConfig.
func (c *Config[T]) LoadAs(path string) (IConfig[T], error) {
	return nil, nil
}

// Implements the SaveAs() function from IConfig.
func (c Config[T]) SaveAs(paths string) error {
	//Ensure at least one marshaler is bound before continuing
	if c.marshalers.Empty() {
		return ErrNoMarshalers
	}

	//TODO: ensure the number of passed paths matches the number of marshalers that have a backing file

	//Loop over the marshalers
	//var data []byte

	/*
		data, err := json.Marshal(c.Data)
		if err != nil {
			return err
		}
		return os.WriteFile(path, data, 0644)
	*/
	return nil
}
