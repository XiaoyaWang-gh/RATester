package udp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestShutdown_2(t *testing.T) {
	type fields struct {
		mu        sync.RWMutex
		conns     map[string]*Conn
		accepting bool
		timeout   time.Duration
	}
	type args struct {
		graceTimeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
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

			l := &Listener{
				mu:        tt.fields.mu,
				conns:     tt.fields.conns,
				accepting: tt.fields.accepting,
				timeout:   tt.fields.timeout,
			}
			if err := l.Shutdown(tt.args.graceTimeout); (err != nil) != tt.wantErr {
				t.Errorf("Listener.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
