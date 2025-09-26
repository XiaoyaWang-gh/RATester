package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestGetDependencyManagerForScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO(bep)
}
