package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_3(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
	}
	tests := []struct {
		name    string
		lp      *Provider
		args    args
		want    bool
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

			got, err := tt.lp.SessionExist(tt.args.ctx, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Provider.SessionExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
