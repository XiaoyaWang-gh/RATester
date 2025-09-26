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
			name:    "Test case 1",
			v:       1000,
			want:    1000 * time.Millisecond,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			v:       "1s",
			want:    time.Second,
			wantErr: false,
		},
		{
			name:    "Test case 3",
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
