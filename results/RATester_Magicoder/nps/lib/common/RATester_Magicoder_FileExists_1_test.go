package common

import (
	"fmt"
	"testing"
)

func TestFileExists_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "File exists",
			args: args{name: "testdata/existing_file.txt"},
			want: true,
		},
		{
			name: "File does not exist",
			args: args{name: "testdata/non_existing_file.txt"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FileExists(tt.args.name); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
