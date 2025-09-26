package resources

import (
	"fmt"
	"testing"
)

func TestReplaceResourcePlaceholders_1(t *testing.T) {
	type args struct {
		in      string
		counter int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				in:      "Hello, {0}",
				counter: 1,
			},
			want: "Hello, 1",
		},
		{
			name: "Test case 2",
			args: args{
				in:      "Goodbye, {0}",
				counter: 2,
			},
			want: "Goodbye, 2",
		},
		{
			name: "Test case 3",
			args: args{
				in:      "{0} is the number",
				counter: 3,
			},
			want: "3 is the number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := replaceResourcePlaceholders(tt.args.in, tt.args.counter); got != tt.want {
				t.Errorf("replaceResourcePlaceholders() = %v, want %v", got, tt.want)
			}
		})
	}
}
