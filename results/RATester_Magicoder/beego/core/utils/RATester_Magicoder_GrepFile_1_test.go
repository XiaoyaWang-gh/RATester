package utils

import (
	"fmt"
	"testing"
)

func TestGrepFile_1(t *testing.T) {
	type args struct {
		patten   string
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
				patten:   "test",
				filename: "test.txt",
			},
			wantErr: false,
		},
		{
			name: "TestGrepFile_1",
			args: args{
				patten:   "test",
				filename: "test1.txt",
			},
			wantErr: true,
		},
		{
			name: "TestGrepFile_2",
			args: args{
				patten:   "test",
				filename: "test2.txt",
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

			_, err := GrepFile(tt.args.patten, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrepFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
