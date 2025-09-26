package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetName_1(t *testing.T) {
	tests := []struct {
		name string
		in   *HelloRequest
		want string
	}{
		{
			name: "Test case 1",
			in:   &HelloRequest{Name: "John Doe"},
			want: "John Doe",
		},
		{
			name: "Test case 2",
			in:   &HelloRequest{Name: "Jane Doe"},
			want: "Jane Doe",
		},
		{
			name: "Test case 3",
			in:   nil,
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

			if got := tt.in.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
