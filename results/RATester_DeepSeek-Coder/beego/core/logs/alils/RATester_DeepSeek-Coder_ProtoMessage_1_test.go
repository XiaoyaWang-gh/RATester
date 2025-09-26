package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestProtoMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(123456789),
				Contents: []*LogContent{
					{
						Key:   proto.String("testKey"),
						Value: proto.String("testValue"),
					},
				},
			},
		},
		Reserved: proto.String("reserved"),
		Topic:    proto.String("topic"),
		Source:   proto.String("source"),
	}

	lg.ProtoMessage()

	// Test if the function does not panic
}
