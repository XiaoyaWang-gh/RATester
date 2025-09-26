package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalString_1(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
		want    *BandwidthQuantity
	}{
		{
			name:    "MB",
			s:       "10MB",
			wantErr: false,
			want:    &BandwidthQuantity{s: "10MB", i: 10 * MB},
		},
		{
			name:    "KB",
			s:       "10KB",
			wantErr: false,
			want:    &BandwidthQuantity{s: "10KB", i: 10 * KB},
		},
		{
			name:    "invalid",
			s:       "10KBa",
			wantErr: true,
			want:    nil,
		},
		{
			name:    "empty",
			s:       "",
			wantErr: false,
			want:    &BandwidthQuantity{s: "", i: 0},
		},
		{
			name:    "space",
			s:       "   ",
			wantErr: false,
			want:    &BandwidthQuantity{s: "", i: 0},
		},
		{
			name:    "unit not support",
			s:       "10B",
			wantErr: true,
			want:    nil,
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
			if err := q.UnmarshalString(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(q, tt.want) {
				t.Errorf("UnmarshalString() = %v, want %v", q, tt.want)
			}
		})
	}
}
