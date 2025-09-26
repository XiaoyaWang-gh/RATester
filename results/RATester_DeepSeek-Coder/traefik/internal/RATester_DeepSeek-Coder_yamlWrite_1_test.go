package main

import (
	"fmt"
	"testing"
)

func TestYamlWrite_1(t *testing.T) {
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
			name: "Test with valid file and element",
			args: args{
				outputFile: "test.yaml",
				element:    map[string]string{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid file",
			args: args{
				outputFile: "/invalid/path/test.yaml",
				element:    map[string]string{"key": "value"},
			},
			wantErr: true,
		},
		{
			name: "Test with nil element",
			args: args{
				outputFile: "test.yaml",
				element:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := yamlWrite(tt.args.outputFile, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
