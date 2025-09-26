package ecs

import (
	"fmt"
	"testing"
)

func TestGetPort_1(t *testing.T) {
	type args struct {
		instance   ecsInstance
		serverPort string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPort(tt.args.instance, tt.args.serverPort); got != tt.want {
				t.Errorf("getPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
