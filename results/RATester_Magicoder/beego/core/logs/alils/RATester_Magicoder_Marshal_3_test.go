package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestMarshal_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logGroup := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(123456789),
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

	data, err := logGroup.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	unmarshaledLogGroup := &LogGroup{}
	if err := unmarshaledLogGroup.Unmarshal(data); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(logGroup, unmarshaledLogGroup) {
		t.Fatalf("Marshal and Unmarshal failed: %v != %v", logGroup, unmarshaledLogGroup)
	}
}
