package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(1234567890),
				Contents: []*LogContent{
					{
						Key:   proto.String("key1"),
						Value: proto.String("value1"),
					},
					{
						Key:   proto.String("key2"),
						Value: proto.String("value2"),
					},
				},
			},
		},
		Reserved: proto.String("reserved"),
		Topic:    proto.String("topic"),
		Source:   proto.String("source"),
	}

	lg.Reset()

	if len(lg.Logs) != 0 {
		t.Errorf("Expected Logs to be empty, but got %v", lg.Logs)
	}
	if lg.Reserved != nil {
		t.Errorf("Expected Reserved to be nil, but got %v", *lg.Reserved)
	}
	if lg.Topic != nil {
		t.Errorf("Expected Topic to be nil, but got %v", *lg.Topic)
	}
	if lg.Source != nil {
		t.Errorf("Expected Source to be nil, but got %v", *lg.Source)
	}
}
