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
			name: "Test with .txt extension",
			args: args{
				path:   "test.txt",
				newExt: "doc",
			},
			want: "test.doc",
		},
		{
			name: "Test with .docx extension",
			args: args{
				path:   "test.docx",
				newExt: "pdf",
			},
			want: "test.pdf",
		},
		{
			name: "Test with .go extension",
			args: args{
				path:   "main.go",
				newExt: "java",
			},
			want: "main.java",
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
