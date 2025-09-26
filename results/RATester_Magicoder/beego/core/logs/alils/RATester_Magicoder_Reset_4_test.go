package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestReset_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Log{
		Time:            proto.Uint32(1),
		Contents:        []*LogContent{&LogContent{Key: proto.String("key"), Value: proto.String("value")}},
		XXXUnrecognized: []byte{},
	}
	l.Reset()
	if l.Time != nil {
		t.Errorf("Expected Time to be nil, but got %v", l.Time)
	}
	if len(l.Contents) != 0 {
		t.Errorf("Expected Contents to be empty, but got %v", l.Contents)
	}
	if len(l.XXXUnrecognized) != 0 {
		t.Errorf("Expected XXXUnrecognized to be empty, but got %v", l.XXXUnrecognized)
	}
}
