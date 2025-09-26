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
				in: "test.txt",
			},
			want:  "test",
			want1: "txt",
		},
		{
			name: "Test case 2",
			args: args{
				in: "test.txt.gz",
			},
			want:  "test.txt",
			want1: "gz",
		},
		{
			name: "Test case 3",
			args: args{
				in: "test",
			},
			want:  "test",
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
