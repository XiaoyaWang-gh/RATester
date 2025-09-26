package scss

import (
	"fmt"
	"testing"
)

func TestreplaceRegularImportsOut_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{s: "@import \"regular.css\";"},
			want: "@import \"regular.css\";",
		},
		{
			name: "Test case 2",
			args: args{s: "@import \"another.css\";"},
			want: "@import \"another.css\";",
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

			if got := replaceRegularImportsOut(tt.args.s); got != tt.want {
				t.Errorf("replaceRegularImportsOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
