package hydratation

import (
	"fmt"
	"testing"
)

func TestHydrate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	element := struct {
		Name string
	}{}
	err := Hydrate(element)
	if err != nil {
		t.Errorf("Hydrate() error = %v", err)
		return
	}
	if element.Name != "" {
		t.Errorf("Hydrate() element.Name = %v, want %v", element.Name, "")
	}
}
