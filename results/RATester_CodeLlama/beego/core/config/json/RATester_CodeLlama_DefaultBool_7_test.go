package json

import (
	"fmt"
	"testing"
)

func TestDefaultBool_7(t *testing.T) {
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
	if v := c.DefaultBool("key", false); v != true {
		t.Fatalf("DefaultBool failed")
	}
	if v := c.DefaultBool("key1", false); v != false {
		t.Fatalf("DefaultBool failed")
	}
}
