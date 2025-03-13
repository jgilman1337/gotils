package marshaler

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type cfg struct {
	Foo    string
	Bar    int
	FooBar map[string]int
	Baz    []string
}

var (
	//The expected default value of a test config object
	dat = cfg{
		Foo:    "hello world",
		Bar:    42,
		FooBar: map[string]int{"foo": 1, "bar": 2, "baz": 3},
		Baz:    []string{"foo", "bar", "baz"},
	}

	//The expected output of a JSON marshal
	jout = []byte(`{"Foo":"hello world","Bar":42,"FooBar":{"bar":2,"baz":3,"foo":1},"Baz":["foo","bar","baz"]}`)

	//A JSON marshaler instance for testing
	mjson = Json{}
)

// Tests the marshaling function of the JSON marshaler struct.
func TestMarshalJson(t *testing.T) {
	//Run the test
	mjson.Minified = true
	actual, err := mjson.Marshal(&dat)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", actual)

	//Check for accuracy
	if !bytes.Equal(actual, jout) {
		t.Fatalf("incorrect JSON marshal output; got `%s`, expected `%s`\n", actual, jout)
	}
}

// Tests the unmarshaling function of the JSON marshaler struct.
func TestUMarshalJson(t *testing.T) {
	//Run the test
	var actual cfg
	if err := mjson.UMarshal(jout, &actual); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	//Check for accuracy
	if !reflect.DeepEqual(actual, dat) {
		t.Fatalf("incorrect JSON unmarshal output; got `%+v`, expected `%+v`\n", actual, dat)
	}
}
