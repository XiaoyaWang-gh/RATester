package label

import (
	"fmt"
	"testing"
)

func TestDecode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	labels := map[string]string{
		"name": "test",
	}
	element := &struct {
		Name string `json:"name"`
	}{}
	filters := []string{"name"}
	err := Decode(labels, element, filters...)
	if err != nil {
		t.Errorf("Decode failed: %v", err)
	}
	if element.Name != "test" {
		t.Errorf("Decode failed: %v", element.Name)
	}
}
