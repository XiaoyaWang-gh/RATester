package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestString_1(t *testing.T) {
	testCases := []struct {
		name string
		msg  *Test_OptionalGroup
		want string
	}{
		{
			name: "Empty",
			msg:  &Test_OptionalGroup{},
			want: "<nil>",
		},
		{
			name: "Populated",
			msg: &Test_OptionalGroup{
				RequiredField: proto.String("test"),
			},
			want: `RequiredField:"test"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.msg.String()
			if got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}
