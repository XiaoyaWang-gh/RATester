package cache

import (
	"fmt"
	"testing"
)

func TestGetInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input interface{}
		want  int
	}

	tests := []test{
		{input: 10, want: 10},
		{input: int32(32), want: 32},
		{input: int64(64), want: 64},
		{input: "70", want: 70},
		{input: "invalid", want: 0},
	}

	for _, tc := range tests {
		got := GetInt(tc.input)
		if got != tc.want {
			t.Errorf("GetInt(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
