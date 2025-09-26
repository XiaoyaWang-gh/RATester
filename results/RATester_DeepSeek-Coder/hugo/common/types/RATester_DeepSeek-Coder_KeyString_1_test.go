package types

import (
	"fmt"
	"testing"
)

func TestKeyString_1(t *testing.T) {
	tests := []struct {
		name string
		k    KeyValues
		want string
	}{
		{
			name: "Test Case 1",
			k: KeyValues{
				Key:    "testKey",
				Values: []any{"value1", "value2"},
			},
			want: "testKey",
		},
		{
			name: "Test Case 2",
			k: KeyValues{
				Key:    123,
				Values: []any{456, 789},
			},
			want: "123",
		},
		{
			name: "Test Case 3",
			k: KeyValues{
				Key:    true,
				Values: []any{false, true},
			},
			want: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.k.KeyString(); got != tt.want {
				t.Errorf("KeyValues.KeyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
