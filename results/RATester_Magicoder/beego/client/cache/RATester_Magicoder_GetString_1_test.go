package cache

import (
	"fmt"
	"testing"
)

func TestGetString_1(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"string", "test", "test"},
		{"[]byte", []byte("test"), "test"},
		{"nil", nil, ""},
		{"int", 123, "123"},
		{"float", 123.456, "123.456"},
		{"bool", true, "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetString(tt.v); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
