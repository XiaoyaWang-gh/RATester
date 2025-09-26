package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_8(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		config      string
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
			if err := pder.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("CookieProvider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
