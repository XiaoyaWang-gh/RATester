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
				Key:    "test",
				Values: []any{1, 2, 3},
			},
			want: "test",
		},
		{
			name: "Test Case 2",
			k: KeyValues{
				Key:    123,
				Values: []any{"a", "b", "c"},
			},
			want: "123",
		},
		{
			name: "Test Case 3",
			k: KeyValues{
				Key:    nil,
				Values: []any{},
			},
			want: "",
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
				t.Errorf("KeyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
