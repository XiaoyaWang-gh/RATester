package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetContents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	log := &Log{
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
	}

	contents := log.GetContents()
	if len(contents) != 2 {
		t.Errorf("Expected 2 contents, got %d", len(contents))
	}

	if contents[0].GetKey() != "key1" || contents[0].GetValue() != "value1" {
		t.Errorf("Expected key1 and value1, got %s and %s", contents[0].GetKey(), contents[0].GetValue())
	}

	if contents[1].GetKey() != "key2" || contents[1].GetValue() != "value2" {
		t.Errorf("Expected key2 and value2, got %s and %s", contents[1].GetKey(), contents[1].GetValue())
	}
}
