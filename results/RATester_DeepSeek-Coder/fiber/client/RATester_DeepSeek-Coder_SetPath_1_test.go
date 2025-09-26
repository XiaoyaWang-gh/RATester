package client

import (
	"fmt"
	"testing"
)

func TestSetPath_1(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		f    *File
		args args
		want string
	}{
		{
			name: "TestSetPath",
			f:    &File{},
			args: args{
				p: "test_path",
			},
			want: "test_path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.f.SetPath(tt.args.p)
			if tt.f.path != tt.want {
				t.Errorf("File.SetPath() = %v, want %v", tt.f.path, tt.want)
			}
		})
	}
}
