package json

import (
	"fmt"
	"testing"
)

func TestBool_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": true,
		},
	}
	b, err := c.Bool("key")
	if err != nil {
		t.Error(err)
	}
	if b != true {
		t.Errorf("expect true, but get %v", b)
	}
}
