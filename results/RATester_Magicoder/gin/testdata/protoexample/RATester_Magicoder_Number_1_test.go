package protoexample

import (
	"fmt"
	"testing"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func TestNumber_1(t *testing.T) {
	tests := []struct {
		name string
		x    FOO
		want protoreflect.EnumNumber
	}{
		{
			name: "Test case 1",
			x:    1,
			want: protoreflect.EnumNumber(1),
		},
		{
			name: "Test case 2",
			x:    2,
			want: protoreflect.EnumNumber(2),
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

			if got := tt.x.Number(); got != tt.want {
				t.Errorf("Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
