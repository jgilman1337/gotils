package cfg

import (
	"fmt"
	"testing"
)

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
	dat = NewConfig(cfgtest{
		Foo:    "hello world",
		Bar:    42,
		FooBar: map[string]int{"foo": 1, "bar": 2, "baz": 3},
		Baz:    []string{"foo", "bar", "baz"},
	})
)

// Tests the "copy defaults" functionality of the configuration utility.
func TestCopyDefault(t *testing.T) {

}

// Tests the "defaults" functionality of the configuration utility.
func TestDefaults(t *testing.T) {
	//Run the test
	actual, err := NewConfig(cfgtest{}).Defaults()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	//Check for accuracy
	if !dat.Equal(actual) {
		t.Fatalf("incorrect defaults output; got `%+v`, expected `%+v`\n", actual.Data(), dat.Data())
	}
}
