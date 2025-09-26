package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestSize_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	log := &Log{
		Time:            proto.Uint32(1),
		Contents:        []*LogContent{&LogContent{Key: proto.String("key"), Value: proto.String("value")}},
		XXXUnrecognized: []byte{},
	}
	expectedSize := 1 + 1 + sovLog(uint64(*log.Time)) + 1 + log.Contents[0].Size() + sovLog(uint64(log.Contents[0].Size())) + len(log.XXXUnrecognized)
	actualSize := log.Size()
	if expectedSize != actualSize {
		t.Errorf("Expected size: %d, Actual size: %d", expectedSize, actualSize)
	}
}
