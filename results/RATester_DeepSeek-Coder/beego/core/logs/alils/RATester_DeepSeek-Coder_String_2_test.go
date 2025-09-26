package alils

import (
	"fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	tests := []struct {
		name string
		m    *LogGroup
		want string
	}{
		{
			name: "Test Case 1",
			m: &LogGroup{
				Logs: []*Log{
					{
						Time: func() *uint32 { v := uint32(1630407680); return &v }(),
						Contents: []*LogContent{
							{
								Key:   func() *string { v := "key"; return &v }(),
								Value: func() *string { v := "value"; return &v }(),
							},
						},
					},
				},
				Reserved: func() *string { v := "reserved"; return &v }(),
				Topic:    func() *string { v := "topic"; return &v }(),
				Source:   func() *string { v := "source"; return &v }(),
			},
			want: "Logs: [{Time:1630407680 Contents: [{Key:\"key\" Value:\"value\"}]}], Reserved:\"reserved\", Topic:\"topic\", Source:\"source\"",
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
				t.Errorf("LogGroup.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
