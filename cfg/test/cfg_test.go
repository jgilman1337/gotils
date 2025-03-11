package test

import (
	"fmt"
	"testing"

	"github.com/jgilman1337/gotils/cfg"
)

func TestTemp(t *testing.T) {
	//Create a new config object backed by a JSON file
	actual, err := cfg.NewWithJson[cfgtest]().Defaults()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", actual)

	actual.SaveAs("test.json") //TODO: NOP (for now)

	/*
		actual, err := cfg.NewConfig(cfgtest{}).Defaults()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", actual)

		//actual.
	*/
}

// Tests the "defaults" functionality of the configuration utility.
func TestDefaults(t *testing.T) {
	//Run the test
	actual, err := cfg.NewConfig(cfgtest{}).Defaults()
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
