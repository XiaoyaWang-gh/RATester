package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestProtoMessage_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logContent := &LogContent{
		Key:   proto.String("testKey"),
		Value: proto.String("testValue"),
	}

	logContent.ProtoMessage()

	// Add assertions here if needed
}
