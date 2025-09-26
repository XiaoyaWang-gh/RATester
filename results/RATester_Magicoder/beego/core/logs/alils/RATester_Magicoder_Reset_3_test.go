package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestReset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Logs: []*Log{
					{
						Time: proto.Uint32(123456789),
						Contents: []*LogContent{
							{
								Key:   proto.String("key"),
								Value: proto.String("value"),
							},
						},
					},
				},
			},
		},
	}

	lg.Reset()

	if len(lg.LogGroups) != 0 {
		t.Errorf("Expected LogGroups to be empty, but got %v", lg.LogGroups)
	}
}
