package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestMarshal_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Logs: []*Log{
					{
						Time: proto.Uint32(1),
					},
				},
			},
		},
	}
	data, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("data is empty")
	}
}
