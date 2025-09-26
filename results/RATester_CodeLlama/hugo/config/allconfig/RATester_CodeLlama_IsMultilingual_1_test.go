package allconfig

import (
	"fmt"
	"testing"
)

func TestIsMultilingual_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.IsMultilingual() {
		t.Error("IsMultilingual() should be false")
	}
}
