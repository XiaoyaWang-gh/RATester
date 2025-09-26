package herrors

import (
	"fmt"
	"testing"
)

func TestextNoDelimiter_2(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{filename: "test.txt"},
			want: "txt",
		},
		{
			name: "Test case 2",
			args: args{filename: "test.go"},
			want: "go",
		},
		{
			name: "Test case 3",
			args: args{filename: "test"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extNoDelimiter(tt.args.filename); got != tt.want {
				t.Errorf("extNoDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
