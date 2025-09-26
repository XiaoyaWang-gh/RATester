package utils

import (
	"fmt"
	"testing"
)

func TestSearchFile_1(t *testing.T) {
	type args struct {
		filename string
		paths    []string
	}
	tests := []struct {
		name         string
		args         args
		wantFullpath string
		wantErr      bool
	}{
		{
			name: "Test case 1",
			args: args{
				filename: "test.txt",
				paths:    []string{"/path1", "/path2"},
			},
			wantFullpath: "/path1/test.txt",
			wantErr:      false,
		},
		{
			name: "Test case 2",
			args: args{
				filename: "test.txt",
				paths:    []string{"/path3", "/path4"},
			},
			wantFullpath: "",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotFullpath, err := SearchFile(tt.args.filename, tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFullpath != tt.wantFullpath {
				t.Errorf("SearchFile() = %v, want %v", gotFullpath, tt.wantFullpath)
			}
		})
	}
}
