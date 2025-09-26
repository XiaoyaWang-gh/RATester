package web

import (
	"fmt"
	"testing"
)

func TestDefaultBool_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	key := "key"
	defaultVal := true
	b.DefaultBool(key, defaultVal)
}
