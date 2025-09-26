package types

import (
	"fmt"
	"testing"
)

func TestString_22(t *testing.T) {
	tests := []struct {
		name string
		k    KeyValues
		want string
	}{
		{
			name: "Test case 1",
			k: KeyValues{
				Key:    "key1",
				Values: []any{"value1", "value2"},
			},
			want: "key1: [value1 value2]",
		},
		{
			name: "Test case 2",
			k: KeyValues{
				Key:    "key2",
				Values: []any{"value3", "value4"},
			},
			want: "key2: [value3 value4]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
