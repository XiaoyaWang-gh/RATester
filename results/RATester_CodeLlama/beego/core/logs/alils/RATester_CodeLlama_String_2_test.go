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

	m := &LogGroup{
		Logs: []*Log{
			{
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
			},
		},
		Reserved: proto.String("reserved"),
	}
	want := `LogGroup{Logs:[]*Log{Log{Time:1234567890,Contents:[]*LogContent{LogContent{Key:"key1",Value:"value1"},LogContent{Key:"key2",Value:"value2"}}}},Reserved:"reserved"}`
	if got := m.String(); got != want {
		t.Errorf("m.String() = %q, want %q", got, want)
	}
}
