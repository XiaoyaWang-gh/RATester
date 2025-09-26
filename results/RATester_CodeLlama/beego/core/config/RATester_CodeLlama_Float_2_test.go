package config

import (
	"fmt"
	"testing"
)

func TestFloat_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &BaseConfiger{}
	key := "a.b.c"
	res, err := c.Float(key)
	if err != nil {
		t.Errorf("Float() error = %v, want nil", err)
		return
	}
	if res != 0 {
		t.Errorf("Float() res = %v, want 0", res)
	}
}
