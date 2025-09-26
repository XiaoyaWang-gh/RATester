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

	logGroup := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(123),
				Contents: []*LogContent{
					{
						Key:   proto.String("key"),
						Value: proto.String("value"),
					},
				},
			},
		},
		Reserved: proto.String("reserved"),
		Topic:    proto.String("topic"),
		Source:   proto.String("source"),
	}

	logGroup.Reset()

	if len(logGroup.Logs) != 0 {
		t.Errorf("Expected Logs to be empty, got %v", logGroup.Logs)
	}

	if logGroup.Reserved != nil {
		t.Errorf("Expected Reserved to be nil, got %v", *logGroup.Reserved)
	}

	if logGroup.Topic != nil {
		t.Errorf("Expected Topic to be nil, got %v", *logGroup.Topic)
	}

	if logGroup.Source != nil {
		t.Errorf("Expected Source to be nil, got %v", *logGroup.Source)
	}
}
