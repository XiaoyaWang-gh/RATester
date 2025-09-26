package gin

import (
	"fmt"
	"testing"
)

func TestStack_1(t *testing.T) {
	tests := []struct {
		name string
		skip int
		want string
	}{
		{
			name: "Test case 1",
			skip: 0,
			want: "test.go:14 (0x49f820)\n\tstack.go:23: func stack(skip int) []byte {\n",
		},
		{
			name: "Test case 2",
			skip: 1,
			want: "test.go:28 (0x49f820)\n\tstack.go:23: func stack(skip int) []byte {\n",
		},
		{
			name: "Test case 3",
			skip: 2,
			want: "test.go:42 (0x49f820)\n\tstack.go:23: func stack(skip int) []byte {\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := stack(tt.skip); string(got) != tt.want {
				t.Errorf("stack() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
