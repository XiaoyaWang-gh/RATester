package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	log := &Log{
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
	}

	expected := `Time:1234567890 Contents:[{Key:"key1" Value:"value1"} {Key:"key2" Value:"value2"}]`
	if got := log.String(); got != expected {
		t.Errorf("log.String() = %q, want %q", got, expected)
	}
}
