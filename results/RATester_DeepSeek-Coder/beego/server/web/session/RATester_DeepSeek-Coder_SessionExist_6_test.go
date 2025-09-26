package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_6(t *testing.T) {
	ctx := context.Background()
	pder := &CookieProvider{}

	tests := []struct {
		name    string
		ctx     context.Context
		sid     string
		want    bool
		wantErr bool
	}{
		{
			name:    "Test case 1",
			ctx:     ctx,
			sid:     "test_session_id",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			ctx:     ctx,
			sid:     "",
			want:    false,
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

			got, err := pder.SessionExist(tt.ctx, tt.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SessionExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
