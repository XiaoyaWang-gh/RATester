package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestProtoMessage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	log := &Log{
		Time: proto.Uint32(1),
		Contents: []*LogContent{
			{
				Key:   proto.String("key"),
				Value: proto.String("value"),
			},
		},
	}

	log.ProtoMessage()

	// Add assertions here to validate the expected behavior
}
