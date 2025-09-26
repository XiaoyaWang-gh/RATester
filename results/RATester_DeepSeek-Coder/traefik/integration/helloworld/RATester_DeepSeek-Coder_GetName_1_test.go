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
			name: "Test with non-nil HelloRequest",
			in:   &HelloRequest{Name: "John"},
			want: "John",
		},
		{
			name: "Test with nil HelloRequest",
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

			got := tt.in.GetName()
			if got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
