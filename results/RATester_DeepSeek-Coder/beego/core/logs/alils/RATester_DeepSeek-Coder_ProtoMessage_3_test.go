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
		Time: proto.Uint32(1234567890),
		Contents: []*LogContent{
			{
				Key:   proto.String("testKey"),
				Value: proto.String("testValue"),
			},
		},
		XXXUnrecognized: []byte{},
	}

	log.ProtoMessage()

	// Check if the log fields are correctly set
	if *log.Time != 1234567890 {
		t.Errorf("Expected Time to be 1234567890, got %d", *log.Time)
	}

	if len(log.Contents) != 1 {
		t.Errorf("Expected 1 Content, got %d", len(log.Contents))
	}

	if *log.Contents[0].Key != "testKey" {
		t.Errorf("Expected Key to be 'testKey', got %s", *log.Contents[0].Key)
	}

	if *log.Contents[0].Value != "testValue" {
		t.Errorf("Expected Value to be 'testValue', got %s", *log.Contents[0].Value)
	}
}
