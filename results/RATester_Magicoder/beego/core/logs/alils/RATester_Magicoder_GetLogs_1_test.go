package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logGroup := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(1609459200),
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
			{
				Time: proto.Uint32(1609459201),
				Contents: []*LogContent{
					{
						Key:   proto.String("key3"),
						Value: proto.String("value3"),
					},
					{
						Key:   proto.String("key4"),
						Value: proto.String("value4"),
					},
				},
			},
		},
	}

	logs := logGroup.GetLogs()
	if len(logs) != 2 {
		t.Errorf("Expected 2 logs, got %d", len(logs))
	}

	for _, log := range logs {
		if log.GetTime() == 0 {
			t.Error("Log time is 0")
		}
		contents := log.GetContents()
		if len(contents) != 2 {
			t.Errorf("Expected 2 contents, got %d", len(contents))
		}
		for _, content := range contents {
			if content.GetKey() == "" || content.GetValue() == "" {
				t.Error("Content key or value is empty")
			}
		}
	}
}
