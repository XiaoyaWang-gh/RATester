package math

import (
	"fmt"
	"math"
	"testing"
)

func TestPi_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := Namespace{}
	result := ns.Pi()
	expected := math.Pi

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
