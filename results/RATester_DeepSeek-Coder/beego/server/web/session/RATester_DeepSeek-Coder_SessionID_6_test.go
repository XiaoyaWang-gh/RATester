package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionID_6(t *testing.T) {
	type fields struct {
		sid    string
		lock   sync.RWMutex
		values map[interface{}]interface{}
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

			fs := &FileSessionStore{
				sid:    tt.fields.sid,
				lock:   tt.fields.lock,
				values: tt.fields.values,
			}
			if got := fs.SessionID(tt.args.ctx); got != tt.want {
				t.Errorf("FileSessionStore.SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}
