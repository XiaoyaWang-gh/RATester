package utils

import (
	"fmt"
	"testing"
)

func TestGrepFile_1(t *testing.T) {
	type args struct {
		pattern  string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestGrepFile_0",
			args: args{
				pattern:  "test",
				filename: "testdata/test.txt",
			},
			wantErr: false,
		},
		{
			name: "TestGrepFile_1",
			args: args{
				pattern:  "non-exist",
				filename: "testdata/test.txt",
			},
			wantErr: false,
		},
		{
			name: "TestGrepFile_2",
			args: args{
				pattern:  "test",
				filename: "non-exist.txt",
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

			_, err := GrepFile(tt.args.pattern, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrepFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
