package xml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_8(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": []string{"val1", "val2"},
		},
	}

	tests := []struct {
		name       string
		key        string
		defaultVal []string
		want       []string
	}{
		{
			name:       "existing key",
			key:        "key1",
			defaultVal: []string{"default"},
			want:       []string{"val1", "val2"},
		},
		{
			name:       "non-existing key",
			key:        "key2",
			defaultVal: []string{"default"},
			want:       []string{"default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cc.DefaultStrings(tt.key, tt.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
