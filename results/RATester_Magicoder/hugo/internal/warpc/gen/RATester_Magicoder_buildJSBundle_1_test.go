package main

import (
	"fmt"
	"testing"
)

func TestbuildJSBundle_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				filename: "test.js",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				filename: "test2.js",
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

			if err := buildJSBundle(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("buildJSBundle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
