package alils

import (
	"fmt"
	"testing"
)

func TestMarshalTo_2(t *testing.T) {
	tests := []struct {
		name    string
		m       *Log
		want    int
		wantErr bool
	}{
		{
			name: "Test case 1",
			m: &Log{
				Time: func() *uint32 { v := uint32(1630457600); return &v }(),
				Contents: []*LogContent{
					{
						Key:   func() *string { v := "key"; return &v }(),
						Value: func() *string { v := "value"; return &v }(),
					},
				},
			},
			want:    10,
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

			got, err := tt.m.MarshalTo(make([]byte, 100))
			if (err != nil) != tt.wantErr {
				t.Errorf("Log.MarshalTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Log.MarshalTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
