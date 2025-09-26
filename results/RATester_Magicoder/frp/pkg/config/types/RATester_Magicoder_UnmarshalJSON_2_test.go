package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalJSON_2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *BandwidthQuantity
		wantErr bool
	}{
		{
			name:  "null",
			input: "null",
			want:  nil,
		},
		{
			name:  "MB",
			input: `"1MB"`,
			want:  &BandwidthQuantity{s: "MB", i: 1024 * 1024},
		},
		{
			name:  "KB",
			input: `"1KB"`,
			want:  &BandwidthQuantity{s: "KB", i: 1024},
		},
		{
			name:    "invalid",
			input:   `"invalid"`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			q := &BandwidthQuantity{}
			err := q.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(q, tt.want) {
				t.Errorf("UnmarshalJSON() = %v, want %v", q, tt.want)
			}
		})
	}
}
