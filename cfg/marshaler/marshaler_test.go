package marshaler

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jgilman1337/gotils/cfg"
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
	dat = cfg.NewConfig(cfgtest{
		Foo:    "hello world",
		Bar:    42,
		FooBar: map[string]int{"foo": 1, "bar": 2, "baz": 3},
		Baz:    []string{"foo", "bar", "baz"},
	})

	//The expected output of a Json marshal
	jout = []byte(`{"Foo":"hello world","Bar":42,"FooBar":{"bar":2,"baz":3,"foo":1},"Baz":["foo","bar","baz"]}`)

	//A Json marshaler instance for testing
	mjson = Json[cfgtest]{}
)

// Tests the marshaling function of the Json marshaler struct.
func TestMarshalJson(t *testing.T) {
	//Run the test
	actual, err := mjson.MFunc(dat)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", actual)

	//Check for accuracy
	if bytes.Compare(actual, jout) != 0 {
		t.Fatalf("incorrect JSON marshal output; got `%s`, expected `%s`\n", actual, jout)
	}
}

// Tests the unmarshaling function of the Json marshaler struct.
func TestUMarshalJson(t *testing.T) {
	//Run the test
	var actual cfg.Config[cfgtest]
	if err := mjson.UFunc(jout, &actual); err != nil { //TODO: add equals function to the interface, which deep compares against the `Data()` field of both interfaces
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	/*
		//Check for accuracy
		if bytes.Compare(actual, jout) != 0 {
			t.Fatalf("incorrect JSON marshal output; got `%s`, expected `%s`\n", actual, jout)
		}
	*/
}
