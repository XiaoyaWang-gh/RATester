package resources

import (
	"fmt"
	"testing"
)

func TestAddPathIdentifier_1(t *testing.T) {
	type args struct {
		inPath     string
		identifier string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid path and identifier",
			args: args{
				inPath:     "/path/to/file.txt",
				identifier: "test",
			},
			want: "/path/to/filetest.txt",
		},
		{
			name: "Test with valid path and identifier with trailing slash",
			args: args{
				inPath:     "/path/to/file.txt/",
				identifier: "test",
			},
			want: "/path/to/filetest.txt/",
		},
		{
			name: "Test with valid path and identifier with multiple dots",
			args: args{
				inPath:     "/path/to/file.with.dots.txt",
				identifier: "test",
			},
			want: "/path/to/file.with.dots.test.txt",
		},
		{
			name: "Test with valid path and identifier with multiple dots and trailing slash",
			args: args{
				inPath:     "/path/to/file.with.dots.txt/",
				identifier: "test",
			},
			want: "/path/to/file.with.dots.test.txt/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &ResourceTransformationCtx{}
			if got := ctx.addPathIdentifier(tt.args.inPath, tt.args.identifier); got != tt.want {
				t.Errorf("addPathIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
