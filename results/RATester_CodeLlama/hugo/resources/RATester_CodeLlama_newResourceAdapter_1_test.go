package resources

import (
	"fmt"
	"testing"
)

func TestNewResourceAdapter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var spec *Spec
	var lazyPublish bool
	var target transformableResource

	// when
	newResourceAdapter(spec, lazyPublish, target)

	// then
	// ...
}
