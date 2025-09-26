package v1alpha1

import (
	"fmt"
	"testing"
)

func TestTLSOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &version{}
	v.TLSOptions()
}
