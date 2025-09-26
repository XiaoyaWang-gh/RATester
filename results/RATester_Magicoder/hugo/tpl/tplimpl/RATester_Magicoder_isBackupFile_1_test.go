package tplimpl

import (
	"fmt"
	"testing"
)

func TestisBackupFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{path: "test.txt"},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{path: "test.txt~"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{path: "test.txt~~~"},
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

			if got := isBackupFile(tt.args.path); got != tt.want {
				t.Errorf("isBackupFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
