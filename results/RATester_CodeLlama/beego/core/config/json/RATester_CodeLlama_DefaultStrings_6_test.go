package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultStrings_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": []string{"value"},
		},
	}
	key := "key"
	defaultVal := []string{"default"}
	assert.Equal(t, []string{"value"}, c.DefaultStrings(key, defaultVal))
}
