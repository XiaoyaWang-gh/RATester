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
			name: "Test case 1",
			args: args{
				outputFile: "test.yaml",
				element: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				outputFile: "test.yaml",
				element: struct {
					Field1 string
					Field2 int
				}{
					Field1: "value1",
					Field2: 123,
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				outputFile: "non-existent-file.yaml",
				element: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
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

			if err := yamlWrite(tt.args.outputFile, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("yamlWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
