package web

import (
	"fmt"
	"testing"
)

func TestDefaultInt_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	key := "key"
	defaultVal := 1
	if v := b.DefaultInt(key, defaultVal); v != defaultVal {
		t.Errorf("DefaultInt() = %v, want %v", v, defaultVal)
	}
}
