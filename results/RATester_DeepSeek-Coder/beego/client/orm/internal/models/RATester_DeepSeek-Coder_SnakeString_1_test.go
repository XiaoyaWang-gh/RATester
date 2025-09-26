package models

import (
	"fmt"
	"testing"
)

func TestSnakeString_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Test1", "HelloWorld", "hello_world"},
		{"Test2", "hello_world", "hello_world"},
		{"Test3", "HELLOWORLD", "hello_world"},
		{"Test4", "HELLO_WORLD", "hello_world"},
		{"Test5", "helloWorld", "hello_world"},
		{"Test6", "hello_World", "hello_world"},
		{"Test7", "helloWorld_", "hello_world_"},
		{"Test8", "hello_World_", "hello_world_"},
		{"Test9", "HELLOWORLD_", "hello_world_"},
		{"Test10", "HELLO_WORLD_", "hello_world_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SnakeString(tt.arg); got != tt.want {
				t.Errorf("SnakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
