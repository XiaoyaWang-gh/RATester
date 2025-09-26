package tls

import (
	"fmt"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	certificates := &Certificates{}
	expected := "certificates"
	actual := certificates.Type()
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
