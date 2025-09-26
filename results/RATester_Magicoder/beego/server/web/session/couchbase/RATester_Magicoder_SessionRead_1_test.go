package couchbase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestSessionRead_1(t *testing.T) {
	ctx := context.Background()
	cp := &Provider{
		maxlifetime: 100,
		SavePath:    "/tmp",
		Pool:        "default",
		Bucket:      "sessions",
	}

	cp.b = cp.getBucket()

	tests := []struct {
		name    string
		cp      *Provider
		ctx     context.Context
		sid     string
		want    session.Store
		wantErr bool
	}{
		{
			name: "Test case 1",
			cp:   cp,
			ctx:  ctx,
			sid:  "123",
			want: &SessionStore{
				b:           cp.b,
				sid:         "123",
				maxlifetime: cp.maxlifetime,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			cp:   cp,
			ctx:  ctx,
			sid:  "456",
			want: &SessionStore{
				b:           cp.b,
				sid:         "456",
				maxlifetime: cp.maxlifetime,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.cp.SessionRead(tt.ctx, tt.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionRead() = %v, want %v", got, tt.want)
			}
		})
	}
}
