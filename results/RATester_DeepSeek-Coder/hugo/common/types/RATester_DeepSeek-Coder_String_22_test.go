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
				Values: []any{1, 2, 3},
			},
			want: "key2: [1 2 3]",
		},
		{
			name: "Test case 3",
			k: KeyValues{
				Key:    "key3",
				Values: []any{true, false},
			},
			want: "key3: [true false]",
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
				t.Errorf("KeyValues.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
