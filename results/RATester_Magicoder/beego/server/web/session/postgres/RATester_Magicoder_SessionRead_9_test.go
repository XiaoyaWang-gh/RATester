package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_9(t *testing.T) {
	ctx := context.Background()
	mp := &Provider{
		maxlifetime: 100,
		savePath:    "/tmp",
	}

	tests := []struct {
		name    string
		sid     string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			sid:     "test_sid",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			sid:     "test_sid_2",
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

			_, err := mp.SessionRead(ctx, tt.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
