package couchbase

import (
	"context"
	"fmt"
	"sync"
	"testing"

	couchbase "github.com/couchbase/go-couchbase"
)

func TestSessionID_1(t *testing.T) {
	type fields struct {
		b           *couchbase.Bucket
		sid         string
		lock        sync.RWMutex
		values      map[interface{}]interface{}
		maxlifetime int64
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			cs := &SessionStore{
				b:           tt.fields.b,
				sid:         tt.fields.sid,
				lock:        tt.fields.lock,
				values:      tt.fields.values,
				maxlifetime: tt.fields.maxlifetime,
			}
			if got := cs.SessionID(tt.args.ctx); got != tt.want {
				t.Errorf("SessionStore.SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}
