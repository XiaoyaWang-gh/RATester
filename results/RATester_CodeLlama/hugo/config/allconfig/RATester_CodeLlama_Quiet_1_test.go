package allconfig

import (
	"fmt"
	"testing"
)

func TestQuiet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.Quiet() != false {
		t.Errorf("Quiet() = %v, want %v", c.Quiet(), false)
	}
}
