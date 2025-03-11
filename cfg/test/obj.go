package test

import "github.com/jgilman1337/gotils/cfg"

// --- BEGIN test cfg obj
type cfgtest struct {
	Foo    string         `default:"hello world"`
	Bar    int            `default:"42"`
	FooBar map[string]int `default:"{\"foo\": 1, \"bar\": 2, \"baz\": 3}"`
	Baz    []string       `default:"[\"foo\", \"bar\", \"baz\"]"`
}

//--- END test cfg obj

var (
	//The expected default value of a test config object
	dat = cfg.NewConfig(cfgtest{
		Foo:    "hello world",
		Bar:    42,
		FooBar: map[string]int{"foo": 1, "bar": 2, "baz": 3},
		Baz:    []string{"foo", "bar", "baz"},
	})
)
