package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestReset_3(t *testing.T) {
	tests := []struct {
		name string
		m    *LogGroupList
		want *LogGroupList
	}{
		{
			name: "Test case 1",
			m: &LogGroupList{
				LogGroups: []*LogGroup{
					{
						Logs: []*Log{
							{
								Time: ptr.Uint32(123456789),
								Contents: []*LogContent{
									{
										Key:   ptr.String("key"),
										Value: ptr.String("value"),
									},
								},
							},
						},
						Reserved: ptr.String("reserved"),
						Topic:    ptr.String("topic"),
						Source:   ptr.String("source"),
					},
				},
			},
			want: &LogGroupList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.Reset()
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("got %v, want %v", tt.m, tt.want)
			}
		})
	}
}
