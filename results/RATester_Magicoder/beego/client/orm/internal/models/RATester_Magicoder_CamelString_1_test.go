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
		{"Test1", "hello_world", "HelloWorld"},
		{"Test2", "hello_world_again", "HelloWorldAgain"},
		{"Test3", "hello_world_again_again", "HelloWorldAgainAgain"},
		{"Test4", "hello", "Hello"},
		{"Test5", "hello_", "Hello"},
		{"Test6", "_hello", "Hello"},
		{"Test7", "_hello_", "Hello"},
		{"Test8", "__hello__", "Hello"},
		{"Test9", "__hello_world__", "HelloWorld"},
		{"Test10", "hello_world__", "HelloWorld"},
		{"Test11", "__hello_world", "HelloWorld"},
		{"Test12", "__hello__world", "HelloWorld"},
		{"Test13", "__hello__world__", "HelloWorld"},
		{"Test14", "__hello__world__again", "HelloWorldAgain"},
		{"Test15", "__hello__world__again__", "HelloWorldAgain"},
		{"Test16", "__hello__world__again__again", "HelloWorldAgainAgain"},
		{"Test17", "__hello__world__again__again__", "HelloWorldAgainAgain"},
		{"Test18", "__hello__world__again__again__again", "HelloWorldAgainAgainAgain"},
		{"Test19", "__hello__world__again__again__again__", "HelloWorldAgainAgainAgain"},
		{"Test20", "__hello__world__again__again__again__again", "HelloWorldAgainAgainAgainAgain"},
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
