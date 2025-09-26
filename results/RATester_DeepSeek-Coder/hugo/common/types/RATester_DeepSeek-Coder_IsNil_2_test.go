package types

import (
	"fmt"
	"testing"
)

func TestIsNil_2(t *testing.T) {
	type testStruct struct{}

	tests := []struct {
		name string
		arg  any
		want bool
	}{
		{"Nil", nil, true},
		{"Non-nil int", 1, false},
		{"Non-nil string", "test", false},
		{"Non-nil struct", testStruct{}, false},
		{"Nil channel", make(chan int), true},
		{"Non-nil channel", make(chan int, 1), false},
		{"Nil func", func() {}, true},
		{"Non-nil func", func() { fmt.Println("test") }, false},
		{"Nil map", make(map[string]int), true},
		{"Non-nil map", map[string]int{"test": 1}, false},
		{"Nil pointer", new(int), true},
		{"Non-nil pointer", new(int), false},
		{"Nil slice", make([]int, 0), true},
		{"Non-nil slice", []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsNil(tt.arg); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
