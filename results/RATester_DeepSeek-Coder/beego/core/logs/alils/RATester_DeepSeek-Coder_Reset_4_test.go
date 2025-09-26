package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestReset_4(t *testing.T) {
	tests := []struct {
		name string
		m    *Log
		want *Log
	}{
		{
			name: "TestReset",
			m: &Log{
				Time:            proto.Uint32(123),
				Contents:        []*LogContent{{}},
				XXXUnrecognized: []byte{0x01, 0x02, 0x03},
			},
			want: &Log{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.Reset()
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Log.Reset() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}
