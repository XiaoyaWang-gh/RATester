package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetType_1(t *testing.T) {
	tests := []struct {
		name string
		x    *Test
		want int32
	}{
		{
			name: "Test with nil x",
			x:    nil,
			want: Default_Test_Type,
		},
		{
			name: "Test with x.Type nil",
			x:    &Test{},
			want: Default_Test_Type,
		},
		{
			name: "Test with x.Type not nil",
			x:    &Test{Type: proto.Int32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.x.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
