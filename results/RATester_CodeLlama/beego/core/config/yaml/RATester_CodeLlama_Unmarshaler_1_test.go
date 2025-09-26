package yaml

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
			},
		},
	}
	var obj map[string]interface{}
	err := c.Unmarshaler("a", &obj)
	if err != nil {
		t.Fatal(err)
	}
	if obj["b"] != "c" {
		t.Fatalf("unexpected obj: %v", obj)
	}
}
