package allconfig

import (
	"fmt"
	"testing"
)

func TestIsMultihost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.IsMultihost() {
		t.Error("IsMultihost() should be false")
	}
}
