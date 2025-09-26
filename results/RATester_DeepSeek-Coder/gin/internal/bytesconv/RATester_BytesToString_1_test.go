package bytesconv

import (
	"fmt"
	"testing"
)

func TestBytesToString_1(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want string
	}{
		{
			name: "empty byte slice",
			b:    []byte{},
			want: "",
		},
		{
			name: "non-empty byte slice",
			b:    []byte("hello"),
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BytesToString(tt.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
