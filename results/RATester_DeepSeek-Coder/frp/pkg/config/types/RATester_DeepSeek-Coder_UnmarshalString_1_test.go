package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalString_1(t *testing.T) {
	testCases := []struct {
		name    string
		s       string
		wantErr bool
		want    BandwidthQuantity
	}{
		{
			name:    "empty string",
			s:       "",
			wantErr: false,
			want:    BandwidthQuantity{s: "", i: 0},
		},
		{
			name:    "MB",
			s:       "10MB",
			wantErr: false,
			want:    BandwidthQuantity{s: "10MB", i: 10 * 1024 * 1024},
		},
		{
			name:    "KB",
			s:       "10KB",
			wantErr: false,
			want:    BandwidthQuantity{s: "10KB", i: 10 * 1024},
		},
		{
			name:    "unit not support",
			s:       "10GB",
			wantErr: true,
			want:    BandwidthQuantity{s: "", i: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			q := &BandwidthQuantity{}
			err := q.UnmarshalString(tc.s)
			if (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalString() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(*q, tc.want) {
				t.Errorf("UnmarshalString() got = %v, want %v", *q, tc.want)
			}
		})
	}
}
