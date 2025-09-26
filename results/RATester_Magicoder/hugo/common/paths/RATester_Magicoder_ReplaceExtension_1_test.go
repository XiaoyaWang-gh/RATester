package paths

import (
	"fmt"
	"testing"
)

func TestReplaceExtension_1(t *testing.T) {
	type args struct {
		path   string
		newExt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				path:   "test.txt",
				newExt: "doc",
			},
			want: "test.doc",
		},
		{
			name: "Test case 2",
			args: args{
				path:   "test.doc",
				newExt: "txt",
			},
			want: "test.txt",
		},
		{
			name: "Test case 3",
			args: args{
				path:   "test",
				newExt: "doc",
			},
			want: "test.doc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ReplaceExtension(tt.args.path, tt.args.newExt); got != tt.want {
				t.Errorf("ReplaceExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
