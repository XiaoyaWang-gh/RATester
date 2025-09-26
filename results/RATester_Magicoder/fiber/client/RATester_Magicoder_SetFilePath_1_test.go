package client

import (
	"fmt"
	"testing"
)

func TestSetFilePath_1(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want SetFileFunc
	}{
		{
			name: "Test case 1",
			args: args{
				p: "test_path",
			},
			want: func(f *File) {
				f.SetPath("test_path")
			},
		},
		{
			name: "Test case 2",
			args: args{
				p: "another_test_path",
			},
			want: func(f *File) {
				f.SetPath("another_test_path")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := SetFilePath(tt.args.p)
			file := &File{}
			got(file)
			if file.path != tt.args.p {
				t.Errorf("SetFilePath() = %v, want %v", file.path, tt.args.p)
			}
		})
	}
}
