package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jgilman1337/gotils/cfg"
	"github.com/jgilman1337/gotils/cfg/marshaler"
)

var (
	//The expected output of a JSON marshal
	jout = []byte(`{"Foo":"hello world","Bar":42,"FooBar":{"bar":2,"baz":3,"foo":1},"Baz":["foo","bar","baz"]}`)

	//A JSON marshaler instance for testing
	mjson = marshaler.Json[cfgtest]{}
)

// Tests the marshaling function of the JSON marshaler struct.
func TestMarshalJson(t *testing.T) {
	//Run the test
	actual, err := mjson.Marshal(dat)
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
	var actual cfg.Config[cfgtest]
	if err := mjson.UMarshal(jout, &actual); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	//Check for accuracy
	if !dat.Equal(&actual) {
		t.Fatalf("incorrect JSON unmarshal output; got `%+v`, expected `%+v`\n", actual.Data(), dat.Data())
	}
}
