package goldmark

import (
	"fmt"
	"testing"
)

func TestsanitizeAnchorNameString_1(t *testing.T) {
	type args struct {
		s      string
		idType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s:      "test string",
				idType: "id",
			},
			want: "test-string",
		},
		{
			name: "Test Case 2",
			args: args{
				s:      "another test string",
				idType: "id",
			},
			want: "another-test-string",
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

			if got := sanitizeAnchorNameString(tt.args.s, tt.args.idType); got != tt.want {
				t.Errorf("sanitizeAnchorNameString() = %v, want %v", got, tt.want)
			}
		})
	}
}
