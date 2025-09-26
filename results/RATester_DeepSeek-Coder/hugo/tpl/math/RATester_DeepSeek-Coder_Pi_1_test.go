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

	ns := &Namespace{}
	expected := math.Pi
	actual := ns.Pi()

	if actual != expected {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}
