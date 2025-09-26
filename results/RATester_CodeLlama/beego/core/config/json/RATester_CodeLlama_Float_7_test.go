package json

import (
	"fmt"
	"testing"
)

func TestFloat_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": float64(1.1),
		},
	}
	v, err := c.Float("key")
	if err != nil {
		t.Error(err)
	}
	if v != float64(1.1) {
		t.Errorf("want %f, but %f", float64(1.1), v)
	}
}
