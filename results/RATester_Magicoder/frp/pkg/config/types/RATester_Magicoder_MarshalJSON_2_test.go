package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalJSON_2(t *testing.T) {
	tests := []struct {
		name    string
		q       *BandwidthQuantity
		want    []byte
		wantErr bool
	}{
		{
			name: "MB",
			q: &BandwidthQuantity{
				s: "MB",
				i: 1024 * 1024,
			},
			want:    []byte("\"MB\""),
			wantErr: false,
		},
		{
			name: "KB",
			q: &BandwidthQuantity{
				s: "KB",
				i: 1024,
			},
			want:    []byte("\"KB\""),
			wantErr: false,
		},
		{
			name: "Invalid",
			q: &BandwidthQuantity{
				s: "Invalid",
				i: 0,
			},
			want:    nil,
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

			got, err := tt.q.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("BandwidthQuantity.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BandwidthQuantity.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
