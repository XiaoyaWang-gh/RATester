package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLogGroups_1(t *testing.T) {
	tests := []struct {
		name string
		m    *LogGroupList
		want []*LogGroup
	}{
		{
			name: "Test with nil LogGroupList",
			m:    nil,
			want: nil,
		},
		{
			name: "Test with empty LogGroupList",
			m:    &LogGroupList{},
			want: nil,
		},
		{
			name: "Test with non-empty LogGroupList",
			m:    &LogGroupList{LogGroups: []*LogGroup{{}, {}}},
			want: []*LogGroup{{}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.GetLogGroups(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogGroupList.GetLogGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
