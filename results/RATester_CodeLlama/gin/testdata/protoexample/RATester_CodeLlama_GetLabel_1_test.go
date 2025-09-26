package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetLabel_1(t *testing.T) {
	tests := []struct {
		name string
		x    *Test
		want string
	}{
		{
			name: "label is not nil",
			x:    &Test{Label: proto.String("label")},
			want: "label",
		},
		{
			name: "label is nil",
			x:    &Test{},
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

			if got := tt.x.GetLabel(); got != tt.want {
				t.Errorf("GetLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
