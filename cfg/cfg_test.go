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
	datjson = []byte(`{"Foo":"hello world","Bar":42,"FooBar":{"bar":2,"baz":3,"foo":1},"Baz":["foo","bar","baz"]}`)
)

// Tests the "defaults" functionality of the configuration utility using the default provider.
func TestCreastyDefaults(t *testing.T) {
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

// Tests the "defaults" functionality of the configuration utility using a user specified provider.
func TestUserDefaults(t *testing.T) {
	//Specify the defaults provider
	dprovider := func() (*cfgtest, error) {
		return dat.Data(), nil
	}

	//Run the test
	cfg := NewConfig(cfgtest{})
	cfg.DFunc = dprovider
	if _, err := cfg.Defaults(); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", cfg)

	//Check for accuracy
	if !dat.Equal(cfg) {
		t.Fatalf("incorrect defaults output; got `%+v`, expected `%+v`\n", cfg.Data(), dat.Data())
	}
}

// Tests the "loader from bytes" functionality and ensures the JSON marshaler loads it properly.
func TestLoadBytesJson(t *testing.T) {
	//Create the config object and bind marshalers
	cfg := NewConfig(cfgtest{}).
		BindMarshaler(marshaler.Json{})

	//Unmarshal the config object
	if _, err := cfg.LoadBytes(datjson); err != nil {
		t.Fatal(err)
	}

	//Ensure the actual and expected values match
	if !dat.Equal(cfg) {
		t.Fatalf("incorrect unmarshaller output; got `%+v`, expected `%+v`\n", cfg.Data(), dat.Data())
	}
}

// Tests the "loader from file" functionality and ensures the JSON marshaler loads it properly.
func TestLoadFileJson(t *testing.T) {
	//Setup a temp file for the test
	tmpfile, err := os.CreateTemp("", "*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write(datjson); err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	//Create the config object and bind marshalers
	cfg := NewConfig(cfgtest{}).
		BindMarshaler(marshaler.NewJson(tmpfile.Name()))

	//Unmarshal the config object
	if _, err := cfg.LoadPath(); err != nil {
		t.Fatal(err)
	}

	//Ensure the actual and expected values match
	if !dat.Equal(cfg) {
		t.Fatalf("incorrect unmarshaller output; got `%+v`, expected `%+v`\n", cfg.Data(), dat.Data())
	}
}

// Tests the "defaults" functionality and ensures the JSON marshaler saves it properly.
func TestSaveJson(t *testing.T) {
	//Configure the config object and create the configurer for the marshaler
	cfg := NewConfigDefaults[cfgtest]()
	mcfgFunc := func(p string) marshaler.Marshaler {
		jsonm := marshaler.NewJson(p)
		jsonm.Minified = true
		return jsonm
	}

	//Run the test
	testMarshal2File(cfg, mcfgFunc, datjson, t)
}

// Common testing backend for the "save marshal" tests.
func testMarshal2File[T any](cfg *Config[T], marshalerCfger func(p string) marshaler.Marshaler, expected []byte, t *testing.T) {
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
	if !bytes.Equal(expected, actual) {
		t.Fatalf("incorrect marshaler output\n  actual:   %s\n  expected: %s\n", string(actual), string(expected))
	}
}
