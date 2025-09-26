package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionUpdate_1(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
	}
	tests := []struct {
		name    string
		pder    *CookieProvider
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pder := &CookieProvider{}
			if err := pder.SessionUpdate(tt.args.ctx, tt.args.sid); (err != nil) != tt.wantErr {
				t.Errorf("CookieProvider.SessionUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
