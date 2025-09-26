package config

import (
	"fmt"
	"testing"
)

func TestSignature_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type S string
	type C string

	ns := &ConfigNamespace[S, C]{}

	sig := ns.Signature()

	if sig != "" {
		t.Errorf("Expected empty signature, got %v", sig)
	}
}
