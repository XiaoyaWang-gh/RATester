package yaml

import (
	"fmt"
	"testing"
)

func TestInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": 1,
		},
	}
	v, err := c.Int("key")
	if err != nil {
		t.Error(err)
	}
	if v != 1 {
		t.Errorf("expect 1 but get %d", v)
	}
}
