package main

import (
	"fmt"
	"testing"
)

func TestTomlWrite_1(t *testing.T) {
	type args struct {
		outputFile string
		element    any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				outputFile: "test.toml",
				element:    struct{ Name string }{Name: "Test"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				outputFile: "test.toml",
				element:    make(chan int),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tomlWrite(tt.args.outputFile, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("tomlWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
