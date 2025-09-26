package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_7(t *testing.T) {
	a := Alpha{}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{
			name: "Test case 1",
			obj:  "abc",
			want: true,
		},
		{
			name: "Test case 2",
			obj:  "abc123",
			want: false,
		},
		{
			name: "Test case 3",
			obj:  "abc123@",
			want: false,
		},
		{
			name: "Test case 4",
			obj:  "abc123@",
			want: false,
		},
		{
			name: "Test case 5",
			obj:  "abc123@",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := a.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
