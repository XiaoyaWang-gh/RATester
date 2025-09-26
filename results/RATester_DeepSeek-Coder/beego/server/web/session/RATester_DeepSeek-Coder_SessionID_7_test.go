package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSessionID_7(t *testing.T) {
	type fields struct {
		sid          string
		timeAccessed time.Time
		value        map[interface{}]interface{}
		lock         sync.RWMutex
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

			st := &MemSessionStore{
				sid:          tt.fields.sid,
				timeAccessed: tt.fields.timeAccessed,
				value:        tt.fields.value,
				lock:         tt.fields.lock,
			}
			if got := st.SessionID(tt.args.ctx); got != tt.want {
				t.Errorf("MemSessionStore.SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}
