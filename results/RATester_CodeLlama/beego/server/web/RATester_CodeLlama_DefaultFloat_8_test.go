package web

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	key := "key"
	defaultVal := float64(1)
	if got := b.DefaultFloat(key, defaultVal); got != defaultVal {
		t.Errorf("DefaultFloat() = %v, want %v", got, defaultVal)
	}
}
