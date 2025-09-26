package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	testCases := []struct {
		name string
		msg  *StreamExampleReply
		want string
	}{
		{
			name: "Empty",
			msg:  &StreamExampleReply{},
			want: "data: \"\"",
		},
		{
			name: "Populated",
			msg:  &StreamExampleReply{Data: "test"},
			want: "data: \"test\"",
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
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
