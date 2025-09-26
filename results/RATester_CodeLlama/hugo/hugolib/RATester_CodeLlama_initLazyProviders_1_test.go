package hugolib

import (
	"fmt"
	"testing"
)

func TestInitLazyProviders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	ps := &pageState{}
	if err := ps.initLazyProviders(); err != nil {
		t.Fatal(err)
	}

	// TODO: verify results
}
