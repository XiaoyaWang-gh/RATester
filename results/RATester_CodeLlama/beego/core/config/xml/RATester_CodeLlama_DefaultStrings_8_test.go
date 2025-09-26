package xml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": []string{"value1", "value2"},
		},
	}
	key := "key"
	defaultVal := []string{"default1", "default2"}
	if v := c.DefaultStrings(key, defaultVal); !reflect.DeepEqual(v, []string{"value1", "value2"}) {
		t.Errorf("DefaultStrings() = %v, want %v", v, []string{"value1", "value2"})
	}
}
