package main

import (
	"fmt"
	"testing"
)

func TestSaveConfig_1(t *testing.T) {
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
			name: "test1",
			args: args{
				host:     "127.0.0.1",
				vkey:     "123456",
				connType: "tcp",
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
