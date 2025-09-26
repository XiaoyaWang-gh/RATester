package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetRequiredField_1(t *testing.T) {
	tests := []struct {
		name string
		x    *Test_OptionalGroup
		want string
	}{
		{
			name: "Test case 1",
			x: &Test_OptionalGroup{
				RequiredField: proto.String("test"),
			},
			want: "test",
		},
		{
			name: "Test case 2",
			x: &Test_OptionalGroup{
				RequiredField: nil,
			},
			want: "",
		},
		{
			name: "Test case 3",
			x:    nil,
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

			if got := tt.x.GetRequiredField(); got != tt.want {
				t.Errorf("GetRequiredField() = %v, want %v", got, tt.want)
			}
		})
	}
}
