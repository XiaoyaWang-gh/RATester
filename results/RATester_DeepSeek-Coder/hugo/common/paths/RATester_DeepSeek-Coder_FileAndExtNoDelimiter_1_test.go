package paths

import (
	"fmt"
	"testing"
)

func TestFileAndExtNoDelimiter_1(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Test case 1",
			args: args{
				in: "file.ext",
			},
			want:  "file",
			want1: "ext",
		},
		{
			name: "Test case 2",
			args: args{
				in: "file.ext.",
			},
			want:  "file",
			want1: "ext.",
		},
		{
			name: "Test case 3",
			args: args{
				in: "file.",
			},
			want:  "file",
			want1: "",
		},
		{
			name: "Test case 4",
			args: args{
				in: ".ext",
			},
			want:  "",
			want1: "ext",
		},
		{
			name: "Test case 5",
			args: args{
				in: "file",
			},
			want:  "file",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := FileAndExtNoDelimiter(tt.args.in)
			if got != tt.want {
				t.Errorf("FileAndExtNoDelimiter() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FileAndExtNoDelimiter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
