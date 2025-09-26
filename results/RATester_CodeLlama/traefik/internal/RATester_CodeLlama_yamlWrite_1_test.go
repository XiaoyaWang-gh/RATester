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
			name: "test case 1",
			args: args{
				outputFile: "test.yaml",
				element:    "test",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				outputFile: "test.yaml",
				element:    "test",
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

			if err := yamlWrite(tt.args.outputFile, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("yamlWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
