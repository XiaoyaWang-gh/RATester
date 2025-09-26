package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigureFromFile_1(t *testing.T) {
	type args struct {
		path   string
		c      any
		strict bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				path:   "test.yaml",
				c:      &struct{}{},
				strict: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				path:   "nonexistent.yaml",
				c:      &struct{}{},
				strict: true,
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := LoadConfigureFromFile(tt.args.path, tt.args.c, tt.args.strict); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigureFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
