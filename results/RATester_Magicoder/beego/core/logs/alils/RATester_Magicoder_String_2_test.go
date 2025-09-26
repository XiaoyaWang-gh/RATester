package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroup{
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
		},
		Reserved: proto.String("reserved"),
		Topic:    proto.String("topic"),
		Source:   proto.String("source"),
	}

	expected := `Logs:[{Time:1609459200 Contents:[{Key:"key1" Value:"value1"} {Key:"key2" Value:"value2"}]}] Reserved:"reserved" Topic:"topic" Source:"source"`
	if got := lg.String(); got != expected {
		t.Errorf("LogGroup.String() = %v, want %v", got, expected)
	}
}
