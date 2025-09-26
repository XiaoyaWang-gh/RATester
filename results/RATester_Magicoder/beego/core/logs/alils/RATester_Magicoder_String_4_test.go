package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestString_4(t *testing.T) {
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

	expected := `logGroups:<logGroups:<logs:<time:123456789 contents:<key:"key" value:"value" > > > > `
	actual := lg.String()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
