package utils

import (
	"fmt"
	"testing"
)

func TestToStr_1(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		args  []int
		want  string
	}{
		{"bool", true, []int{}, "true"},
		{"float32", float32(3.14), []int{}, "3.14"},
		{"float64", float64(3.14), []int{}, "3.14"},
		{"int", int(123), []int{}, "123"},
		{"int8", int8(123), []int{}, "123"},
		{"int16", int16(123), []int{}, "123"},
		{"int32", int32(123), []int{}, "123"},
		{"int64", int64(123), []int{}, "123"},
		{"uint", uint(123), []int{}, "123"},
		{"uint8", uint8(123), []int{}, "123"},
		{"uint16", uint16(123), []int{}, "123"},
		{"uint32", uint32(123), []int{}, "123"},
		{"uint64", uint64(123), []int{}, "123"},
		{"string", "hello", []int{}, "hello"},
		{"[]byte", []byte("hello"), []int{}, "hello"},
		{"default", struct{}{}, []int{}, "{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToStr(tt.value, tt.args...); got != tt.want {
				t.Errorf("ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
