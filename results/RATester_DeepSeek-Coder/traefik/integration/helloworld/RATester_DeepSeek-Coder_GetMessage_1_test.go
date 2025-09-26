package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetMessage_1(t *testing.T) {
	tests := []struct {
		name string
		m    *HelloReply
		want string
	}{
		{
			name: "Test with non-nil HelloReply",
			m:    &HelloReply{Message: "Hello, World!"},
			want: "Hello, World1",
		},
		{
			name: "Test with nil HelloReply",
			m:    nil,
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

			if got := tt.m.GetMessage(); got != tt.want {
				t.Errorf("HelloReply.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
