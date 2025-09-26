package cache

import (
	"fmt"
	"testing"
)

func TestGetFloat64_1(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want float64
	}{
		{
			name: "float64",
			v:    float64(1.23),
			want: 1.23,
		},
		{
			name: "string",
			v:    "1.23",
			want: 1.23,
		},
		{
			name: "invalid string",
			v:    "abc",
			want: 0,
		},
		{
			name: "nil",
			v:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetFloat64(tt.v); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
