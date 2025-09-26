package models

import (
	"fmt"
	"testing"
)

func TestCamelString_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test1", "hello_world", "helloWorld"},
		{"Test2", "hello_world_again", "helloWorldAgain"},
		{"Test3", "hello_world_again_again", "helloWorldAgainAgain"},
		{"Test4", "hello_world_again_again_again", "helloWorldAgainAgainAgain"},
		{"Test5", "hello_world_again_again_again_again", "helloWorldAgainAgainAgainAgain"},
		{"Test6", "hello_world_again_again_again_again_again", "helloWorldAgainAgainAgainAgainAgain"},
		{"Test7", "hello_world_again_again_again_again_again_again", "helloWorldAgainAgainAgainAgainAgainAgain"},
		{"Test8", "hello_world_again_again_again_again_again_again_again", "helloWorldAgainAgainAgainAgainAgainAgainAgain"},
		{"Test9", "hello_world_again_again_again_again_again_again_again_again", "helloWorldAgainAgainAgainAgainAgainAgainAgainAgain"},
		{"Test10", "hello_world_again_again_again_again_again_again_again_again_again", "helloWorldAgainAgainAgainAgainAgainAgainAgainAgainAgain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CamelString(tt.args); got != tt.want {
				t.Errorf("CamelString() = %v, want %v", got, tt.want)
			}
		})
	}
}
