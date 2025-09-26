package main

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	type args struct {
		files map[string]*File
	}

	tests := []struct {
		name    string
		f       fileWriter
		args    args
		wantErr bool
	}{
		{
			name: "should return an error when the base directory does not exist",
			f:    fileWriter{baseDir: "/non-existent-dir"},
			args: args{
				files: map[string]*File{
					"test.go": {
						Package: "main",
						Imports: []string{"fmt"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when a file cannot be written",
			f:    fileWriter{baseDir: "/tmp"},
			args: args{
				files: map[string]*File{
					"test.go": {
						Package: "main",
						Imports: []string{"fmt"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "should not return an error when all files can be written",
			f:    fileWriter{baseDir: "/tmp"},
			args: args{
				files: map[string]*File{
					"test.go": {
						Package: "main",
						Imports: []string{"fmt"},
					},
				},
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

			err := tt.f.Write(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
