package marshaler

import (
	"bytes"
	"fmt"
	"os"
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

	//The expected output of a TOML marshal (encoded as raw bytes to avoid using multiline strings)
	tout = []byte{0x46, 0x6f, 0x6f, 0x20, 0x3d, 0x20, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x0a, 0x42, 0x61, 0x72, 0x20, 0x3d, 0x20, 0x34, 0x32, 0x0a, 0x42, 0x61, 0x7a, 0x20, 0x3d, 0x20, 0x5b, 0x22, 0x66, 0x6f, 0x6f, 0x22, 0x2c, 0x20, 0x22, 0x62, 0x61, 0x72, 0x22, 0x2c, 0x20, 0x22, 0x62, 0x61, 0x7a, 0x22, 0x5d, 0x0a, 0x0a, 0x5b, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x5d, 0x0a, 0x09, 0x62, 0x61, 0x72, 0x20, 0x3d, 0x20, 0x32, 0x0a, 0x09, 0x62, 0x61, 0x7a, 0x20, 0x3d, 0x20, 0x33, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x20, 0x3d, 0x20, 0x31, 0x0a}

	//The expected output of a YAML marshal (encoded as raw bytes to avoid using multiline strings)
	yout = []byte{0x66, 0x6f, 0x6f, 0x3a, 0x20, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a, 0x62, 0x61, 0x72, 0x3a, 0x20, 0x34, 0x32, 0x0a, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x3a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x72, 0x3a, 0x20, 0x32, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x7a, 0x3a, 0x20, 0x33, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6f, 0x3a, 0x20, 0x31, 0x0a, 0x62, 0x61, 0x7a, 0x3a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x66, 0x6f, 0x6f, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x62, 0x61, 0x72, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x62, 0x61, 0x7a, 0x0a}

	//A JSON marshaler instance for testing
	mjson = Json{}

	//A TOML marshaler instance for testing
	mtoml = Toml{}

	//A YAML marshaler instance for testing
	myaml = Yaml{}
)

// Tests the marshaling function of the JSON marshaler struct.
func TestMarshalJson(t *testing.T) {
	//cfgf := func(m Marshaler) Marshaler {
	cfgf := func(m *Marshaler) {
		j, _ := (*m).(Json)
		j.Minified = true
		*m = &j
	}
	testMarshal(t, dat, jout, mjson, cfgf)
}

// Tests the marshaling function of the TOML marshaler struct.
func TestMarshalToml(t *testing.T) {
	testMarshal(t, dat, tout, mtoml, nil)
}

// Tests the marshaling function of the YAML marshaler struct.
func TestMarshalYaml(t *testing.T) {
	testMarshal(t, dat, yout, myaml, nil)
}

// Tests the unmarshaling function of the JSON marshaler struct.
func TestUMarshalJson(t *testing.T) {
	testUMarshal(t, jout, dat, mjson, nil)
}

// Tests the unmarshaling function of the TOML marshaler struct.
func TestUMarshalToml(t *testing.T) {
	testUMarshal(t, tout, dat, mtoml, nil)
}

// Tests the unmarshaling function of the YAML marshaler struct.
func TestUMarshalYaml(t *testing.T) {
	testUMarshal(t, yout, dat, myaml, nil)
}

// Backend for marshaling tests.
func testMarshal(t *testing.T, in cfg, eout []byte, m Marshaler, cfgf func(m *Marshaler)) {
	//Do additional config if necessary
	if cfgf != nil {
		cfgf(&m)
	}

	//Run the test
	actual, err := m.Marshal(&in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", actual)

	//Check for accuracy
	if !bytes.Equal(actual, eout) {
		t.Fatalf("incorrect %T marshal output; got `%s`, expected `%s`\n", m, actual, eout)
	}
}

// Backend for unmarshaling tests.
func testUMarshal(t *testing.T, in []byte, eout cfg, m Marshaler, cfgf func(m *Marshaler)) {
	//Do additional config if necessary
	if cfgf != nil {
		cfgf(&m)
	}

	//Run the test
	var actual cfg
	if err := m.UMarshal(in, &actual); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	//Check for accuracy
	if !reflect.DeepEqual(actual, eout) {
		t.Fatalf("incorrect %T unmarshal output; got `%+v`, expected `%+v`\n", m, actual, eout)
	}
}

// Dumping utility for marshaled data when creating new marshalers.
func dumpMarshaled(t *testing.T, m Marshaler, d any) {
	actual, err := m.Marshal(&d)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile, err := os.CreateTemp("", "*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write(actual); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("dumped to %s\n", tmpfile.Name())
}
