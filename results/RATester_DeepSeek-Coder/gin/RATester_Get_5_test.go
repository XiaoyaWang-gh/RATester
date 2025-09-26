package gin

import (
	"fmt"
	"testing"
)

func TestGet_5(t *testing.T) {
	params := Params{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	}

	tests := []struct {
		name     string
		params   Params
		key      string
		want     string
		wantBool bool
	}{
		{
			name:     "Existing key",
			params:   params,
			key:      "key2",
			want:     "value2",
			wantBool: true,
		},
		{
			name:     "Non-existing key",
			params:   params,
			key:      "key4",
			want:     "",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, ok := tt.params.Get(tt.key)
			if got != tt.want || ok != tt.wantBool {
				t.Errorf("Params.Get() = %v, %v; want %v, %v", got, ok, tt.want, tt.wantBool)
			}
		})
	}
}
