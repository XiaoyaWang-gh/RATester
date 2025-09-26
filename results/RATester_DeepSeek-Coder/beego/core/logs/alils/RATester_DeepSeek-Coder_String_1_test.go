package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		m    *Log
		want string
	}{
		{
			name: "Test case 1",
			m: &Log{
				Time:            proto.Uint32(1630444800),
				Contents:        []*LogContent{&LogContent{Key: proto.String("key"), Value: proto.String("value")}},
				XXXUnrecognized: []byte{},
			},
			want: "Time:1630444800 Contents:Key:\"key\" Value:\"value\"",
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

			if got := tt.m.String(); got != tt.want {
				t.Errorf("Log.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
