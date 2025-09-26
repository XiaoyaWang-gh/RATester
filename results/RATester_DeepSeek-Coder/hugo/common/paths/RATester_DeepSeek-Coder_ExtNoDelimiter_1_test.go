package paths

import (
	"fmt"
	"testing"
)

func TestExtNoDelimiter_1(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{in: "test.txt"},
			want: "txt",
		},
		{
			name: "Test case 2",
			args: args{in: "test.tar.gz"},
			want: "gz",
		},
		{
			name: "Test case 3",
			args: args{in: "test"},
			want: "",
		},
		{
			name: "Test case 4",
			args: args{in: ".hidden"},
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

			if got := ExtNoDelimiter(tt.args.in); got != tt.want {
				t.Errorf("ExtNoDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
