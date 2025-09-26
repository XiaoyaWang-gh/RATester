package config

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_4(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				filename: "test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				filename: "test2.txt",
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

			if err := SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
