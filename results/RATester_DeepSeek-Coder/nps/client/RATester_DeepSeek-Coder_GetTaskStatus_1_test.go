package client

import (
	"fmt"
	"testing"
)

func TestGetTaskStatus_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				path: "test_config.json",
			},
		},
		{
			name: "Test case 2",
			args: args{
				path: "test_config_not_exist.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			GetTaskStatus(tt.args.path)
		})
	}
}
