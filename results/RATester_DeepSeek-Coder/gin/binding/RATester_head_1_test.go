package binding

import (
	"fmt"
	"testing"
)

func TestHead_1(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Test Case 1",
			args: args{
				str: "Hello, World",
				sep: ",",
			},
			want:  "Hello",
			want1: " World",
		},
		{
			name: "Test Case 2",
			args: args{
				str: "Hello, World",
				sep: "|",
			},
			want:  "Hello, World",
			want1: "",
		},
		{
			name: "Test Case 3",
			args: args{
				str: "",
				sep: ",",
			},
			want:  "",
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

			got, got1 := head(tt.args.str, tt.args.sep)
			if got != tt.want {
				t.Errorf("head() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("head() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
