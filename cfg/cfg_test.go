package cfg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jgilman1337/gotils/cfg/marshaler"
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

// Tests the "defaults" functionality of the configuration utility.
func TestDefaults(t *testing.T) {
	//Run the test
	actual, err := NewConfig(cfgtest{}).Defaults()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	/*
		if err := actual.SaveAs("test.json"); err != nil {
			t.Fatal(err)
		}
	*/

	//Check for accuracy
	if !dat.Equal(actual) {
		t.Fatalf("incorrect defaults output; got `%+v`, expected `%+v`\n", actual.Data(), dat.Data())
	}
}

// Tests the "defaults" functionality and ensures the JSON marshaler saves it properly.
func TestSaveJson(t *testing.T) {
	//Set the expected output
	expected := `{"Foo":"hello world","Bar":42,"FooBar":{"bar":2,"baz":3,"foo":1},"Baz":["foo","bar","baz"]}`

	//Configure the config object and create the configurer for the marshaler
	cfg := NewConfigDefaults[cfgtest]()
	mcfgFunc := func(p string) marshaler.Marshaler {
		jsonm := marshaler.NewJson(p)
		jsonm.Minified = true
		return jsonm
	}

	//Run the test
	testMarshal2File(cfg, mcfgFunc, expected, t)
}

// Common testing backend for the "save marshal" tests.
func testMarshal2File[T any](cfg *Config[T], marshalerCfger func(p string) marshaler.Marshaler, expected string, t *testing.T) {
	//Setup a temp file for the test
	tmpfile, err := os.CreateTemp("", "*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())
	//fmt.Printf("using tempfile %s\n", tmpfile.Name())

	//Attach the marshaler to the config object
	cfg.BindMarshaler(marshalerCfger(tmpfile.Name()))

	//Test to see if the save was successful
	if err := cfg.Save(); err != nil {
		t.Fatal(err)
	}
	actual, err := io.ReadAll(tmpfile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal([]byte(expected), actual) {
		t.Fatalf("incorrect marshaler output\n  actual:   %s\n  expected: %s\n", string(actual), string(expected))
	}
}
