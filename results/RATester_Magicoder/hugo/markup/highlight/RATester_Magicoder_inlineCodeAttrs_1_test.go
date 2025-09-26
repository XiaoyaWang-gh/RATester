package highlight

import (
	"fmt"
	"testing"
)

func TestinlineCodeAttrs_1(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				lang: "go",
			},
			want: ` class="code-inline language-go"`,
		},
		{
			name: "Test case 2",
			args: args{
				lang: "python",
			},
			want: ` class="code-inline language-python"`,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := inlineCodeAttrs(tt.args.lang); got != tt.want {
				t.Errorf("inlineCodeAttrs() = %v, want %v", got, tt.want)
			}
		})
	}
}
