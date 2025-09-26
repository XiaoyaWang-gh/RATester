package ledis

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestSessionRead_3(t *testing.T) {
	ctx := context.Background()
	lp := &Provider{
		maxlifetime: 100,
		SavePath:    "/tmp",
		Db:          0,
	}

	tests := []struct {
		name    string
		ctx     context.Context
		sid     string
		want    session.Store
		wantErr bool
	}{
		{
			name: "Test case 1",
			ctx:  ctx,
			sid:  "test_sid",
			want: &SessionStore{
				sid:         "test_sid",
				values:      make(map[interface{}]interface{}),
				maxlifetime: 100,
			},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			ctx:     ctx,
			sid:     "test_sid_2",
			want:    nil,
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

			got, err := lp.SessionRead(tt.ctx, tt.sid)
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
