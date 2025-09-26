package yaml

import (
	"fmt"
	"testing"
)

func TestFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": 1.2,
		},
	}
	v, err := c.Float("key")
	if err != nil {
		t.Error(err)
	}
	if v != 1.2 {
		t.Errorf("expect 1.2 but get %v", v)
	}
}
