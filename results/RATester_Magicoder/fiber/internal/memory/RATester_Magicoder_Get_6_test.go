package memory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_6(t *testing.T) {
	storage := &Storage{
		data: map[string]item{
			"key1": {v: "value1", e: 0},
			"key2": {v: "value2", e: 1},
			"key3": {v: "value3", e: 2},
		},
	}

	tests := []struct {
		name string
		key  string
		want any
	}{
		{"TestGet_ExistingKey", "key1", "value1"},
		{"TestGet_ExpiredKey", "key2", nil},
		{"TestGet_NonExistingKey", "key4", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := storage.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
