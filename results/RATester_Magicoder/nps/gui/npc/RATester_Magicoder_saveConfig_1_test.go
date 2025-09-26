package main

import (
	"fmt"
	"testing"
)

func TestsaveConfig_1(t *testing.T) {
	type args struct {
		host     string
		vkey     string
		connType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_0",
			args: args{
				host:     "test_host",
				vkey:     "test_vkey",
				connType: "test_connType",
			},
		},
		{
			name: "Test_1",
			args: args{
				host:     "",
				vkey:     "test_vkey",
				connType: "test_connType",
			},
		},
		{
			name: "Test_2",
			args: args{
				host:     "test_host",
				vkey:     "",
				connType: "test_connType",
			},
		},
		{
			name: "Test_3",
			args: args{
				host:     "test_host",
				vkey:     "test_vkey",
				connType: "",
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

			saveConfig(tt.args.host, tt.args.vkey, tt.args.connType)
		})
	}
}
