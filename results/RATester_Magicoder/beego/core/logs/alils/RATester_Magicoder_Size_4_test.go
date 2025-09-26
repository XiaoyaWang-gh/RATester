package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestSize_4(t *testing.T) {
	tests := []struct {
		name string
		m    *LogContent
		want int
	}{
		{
			name: "Test case 1",
			m: &LogContent{
				Key:   proto.String("testKey"),
				Value: proto.String("testValue"),
			},
			want: 10,
		},
		{
			name: "Test case 2",
			m: &LogContent{
				Key:   proto.String(""),
				Value: proto.String(""),
			},
			want: 0,
		},
		{
			name: "Test case 3",
			m: &LogContent{
				Key:   proto.String("longKey"),
				Value: proto.String("longValue"),
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.Size(); got != tt.want {
				t.Errorf("LogContent.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
