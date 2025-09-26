package types

import (
	"fmt"
	"testing"
	"time"
)

func TestToDurationE_1(t *testing.T) {
	tests := []struct {
		name    string
		v       any
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "Test with positive integer",
			v:       100,
			want:    100 * time.Millisecond,
			wantErr: false,
		},
		{
			name:    "Test with negative integer",
			v:       -100,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with string duration",
			v:       "1s",
			want:    time.Second,
			wantErr: false,
		},
		{
			name:    "Test with invalid string",
			v:       "invalid",
			want:    0,
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

			got, err := ToDurationE(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDurationE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToDurationE() = %v, want %v", got, tt.want)
			}
		})
	}
}
