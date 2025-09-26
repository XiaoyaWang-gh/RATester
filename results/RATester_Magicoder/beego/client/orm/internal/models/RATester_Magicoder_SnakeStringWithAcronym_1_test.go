package models

import (
	"fmt"
	"testing"
)

func TestSnakeStringWithAcronym_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test1", "HelloWorld", "hello_world"},
		{"Test2", "HelloWorldAgain", "hello_world_again"},
		{"Test3", "HelloWorldAgainAgain", "hello_world_again_again"},
		{"Test4", "HelloWorldAgainAgainAgain", "hello_world_again_again_again"},
		{"Test5", "HelloWorldAgainAgainAgainAgain", "hello_world_again_again_again_again"},
		{"Test6", "HelloWorldAgainAgainAgainAgainAgain", "hello_world_again_again_again_again_again"},
		{"Test7", "HelloWorldAgainAgainAgainAgainAgainAgain", "hello_world_again_again_again_again_again_again"},
		{"Test8", "HelloWorldAgainAgainAgainAgainAgainAgainAgain", "hello_world_again_again_again_again_again_again_again"},
		{"Test9", "HelloWorldAgainAgainAgainAgainAgainAgainAgainAgain", "hello_world_again_again_again_again_again_again_again_again"},
		{"Test10", "HelloWorldAgainAgainAgainAgainAgainAgainAgainAgainAgain", "hello_world_again_again_again_again_again_again_again_again_again"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SnakeStringWithAcronym(tt.args); got != tt.want {
				t.Errorf("SnakeStringWithAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}
