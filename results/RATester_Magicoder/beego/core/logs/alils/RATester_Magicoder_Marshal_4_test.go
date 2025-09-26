package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestMarshal_4(t *testing.T) {
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
						Time: proto.Uint32(1609459200),
						Contents: []*LogContent{
							{
								Key:   proto.String("test_key"),
								Value: proto.String("test_value"),
							},
						},
					},
				},
			},
		},
	}

	data, err := lg.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	lg2 := &LogGroupList{}
	if err := lg2.Unmarshal(data); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(lg, lg2) {
		t.Fatalf("Marshal and Unmarshal failed: %v != %v", lg, lg2)
	}
}
