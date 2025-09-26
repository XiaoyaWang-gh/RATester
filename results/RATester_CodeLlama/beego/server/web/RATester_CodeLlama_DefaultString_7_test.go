package web

import (
	"fmt"
	"testing"
)

func TestDefaultString_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	key := "key"
	defaultVal := "defaultVal"
	b.DefaultString(key, defaultVal)
}
