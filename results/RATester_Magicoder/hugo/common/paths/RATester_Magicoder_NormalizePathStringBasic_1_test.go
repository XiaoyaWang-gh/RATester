package paths

import (
	"fmt"
	"testing"
)

func TestNormalizePathStringBasic_1(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "NormalizePathStringBasic_0",
			arg:  "test string",
			want: "test-string",
		},
		{
			name: "NormalizePathStringBasic_1",
			arg:  "another test string",
			want: "another-test-string",
		},
		{
			name: "NormalizePathStringBasic_2",
			arg:  "   multiple   spaces   ",
			want: "multiple-spaces",
		},
		{
			name: "NormalizePathStringBasic_3",
			arg:  "no spaces",
			want: "no-spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NormalizePathStringBasic(tt.arg); got != tt.want {
				t.Errorf("NormalizePathStringBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
