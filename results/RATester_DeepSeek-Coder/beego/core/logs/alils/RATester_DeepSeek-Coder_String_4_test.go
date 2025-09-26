package alils

import (
	"fmt"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestString_4(t *testing.T) {
	tests := []struct {
		name string
		m    *LogGroupList
		want string
	}{
		{
			name: "Test case 1",
			m: &LogGroupList{
				LogGroups: []*LogGroup{
					{
						Logs: []*Log{
							{
								Time: ptr.Uint32(1609459200),
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
			want: "logGroups:<logs:<time:1609459200 contents:<key:\"key\" value:\"value\" > > reserved:\"reserved\" topic:\"topic\" source:\"source\" > ",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.String(); got != tt.want {
				t.Errorf("LogGroupList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
