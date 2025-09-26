package ledis

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestSessionRead_3(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
	}
	tests := []struct {
		name    string
		lp      *Provider
		args    args
		want    session.Store
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

			got, err := tt.lp.SessionRead(tt.args.ctx, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.SessionRead() = %v, want %v", got, tt.want)
			}
		})
	}
}
