package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestMarshal_4(t *testing.T) {
	tests := []struct {
		name    string
		lg      *LogGroupList
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			lg: &LogGroupList{
				LogGroups: []*LogGroup{
					{
						Logs: []*Log{
							{
								Time: ptr.Uint32(1620000000),
								Contents: []*LogContent{
									{
										Key:   ptr.String("key1"),
										Value: ptr.String("value1"),
									},
								},
							},
						},
						Reserved: ptr.String("reserved1"),
						Topic:    ptr.String("topic1"),
						Source:   ptr.String("source1"),
					},
				},
			},
			want:    []byte{}, // expected marshaled data
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.lg.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("LogGroupList.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogGroupList.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
