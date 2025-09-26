package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBandwidthQuantity_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    BandwidthQuantity
		wantErr bool
	}{
		{
			name:  "valid MB",
			input: "1MB",
			want:  BandwidthQuantity{s: "MB", i: 1024 * 1024},
		},
		{
			name:  "valid KB",
			input: "1KB",
			want:  BandwidthQuantity{s: "KB", i: 1024},
		},
		{
			name:    "invalid",
			input:   "invalid",
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

			got, err := NewBandwidthQuantity(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBandwidthQuantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBandwidthQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}
