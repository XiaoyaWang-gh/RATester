package memcache

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestSessionRegenerate_13(t *testing.T) {
	type args struct {
		ctx    context.Context
		oldsid string
		sid    string
	}
	tests := []struct {
		name    string
		rp      *MemProvider
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

			got, err := tt.rp.SessionRegenerate(tt.args.ctx, tt.args.oldsid, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemProvider.SessionRegenerate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemProvider.SessionRegenerate() = %v, want %v", got, tt.want)
			}
		})
	}
}
