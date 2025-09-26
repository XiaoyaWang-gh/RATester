package deps

import (
	"fmt"
	"testing"
)

func TestCompile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	var d *Deps
	var prototype *Deps

	// CONTEXT_0
	d = &Deps{}

	// CONTEXT_1
	prototype = &Deps{}

	// CONTEXT_2
	prototype = &Deps{}

	// TEST CASE
	err = d.Compile(prototype)
	if err != nil {
		t.Fatal(err)
	}
}
