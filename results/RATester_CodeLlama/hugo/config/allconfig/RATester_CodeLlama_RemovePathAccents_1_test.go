package allconfig

import (
	"fmt"
	"testing"
)

func TestRemovePathAccents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.RemovePathAccents() {
		t.Error("RemovePathAccents() should be false")
	}
}
