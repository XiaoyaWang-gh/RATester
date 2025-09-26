package paths

import (
	"fmt"
	"testing"
)

func TestHasExt_1(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{p: "test.txt"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{p: "test"},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{p: "test/path"},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{p: "test.path/file"},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{p: "test.path.file"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasExt(tt.args.p); got != tt.want {
				t.Errorf("HasExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
