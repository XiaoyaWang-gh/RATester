package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewKeyValuesStrings_1(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		values []string
		want   KeyValues
	}{
		{
			name:   "Test case 1",
			key:    "key1",
			values: []string{"value1", "value2"},
			want: KeyValues{
				Key:    "key1",
				Values: []any{"value1", "value2"},
			},
		},
		{
			name:   "Test case 2",
			key:    "key2",
			values: []string{"value3", "value4"},
			want: KeyValues{
				Key:    "key2",
				Values: []any{"value3", "value4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewKeyValuesStrings(tt.key, tt.values...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyValuesStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
