package web

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	key := "key"
	defaultVal := int64(1)
	b.DefaultInt64(key, defaultVal)
}
