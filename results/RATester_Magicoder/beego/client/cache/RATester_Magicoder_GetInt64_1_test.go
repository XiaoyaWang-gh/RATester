package cache

import (
	"fmt"
	"testing"
)

func TestGetInt64_1(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want int64
	}{
		{
			name: "int",
			v:    10,
			want: 10,
		},
		{
			name: "int32",
			v:    int32(10),
			want: 10,
		},
		{
			name: "int64",
			v:    int64(10),
			want: 10,
		},
		{
			name: "string",
			v:    "10",
			want: 10,
		},
		{
			name: "default",
			v:    "default",
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

			if got := GetInt64(tt.v); got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
