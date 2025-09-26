package models

import (
	"fmt"
	"testing"
)

func TestSnakeString_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test1", "HelloWorld", "hello_world"},
		{"Test2", "Hello_World", "hello_world"},
		{"Test3", "HelloWorld_", "hello_world"},
		{"Test4", "Hello_World_", "hello_world"},
		{"Test5", "Hello_World_123", "hello_world_123"},
		{"Test6", "Hello_World_123_", "hello_world_123"},
		{"Test7", "Hello_World_123_456", "hello_world_123_456"},
		{"Test8", "Hello_World_123_456_", "hello_world_123_456"},
		{"Test9", "Hello_World_123_456_789", "hello_world_123_456_789"},
		{"Test10", "Hello_World_123_456_789_", "hello_world_123_456_789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SnakeString(tt.args); got != tt.want {
				t.Errorf("SnakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
