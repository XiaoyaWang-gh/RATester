package mirror

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestInc_1(t *testing.T) {
	type fields struct {
		handler          http.Handler
		mirrorHandlers   []*mirrorHandler
		rw               http.ResponseWriter
		routinePool      *safe.Pool
		mirrorBody       bool
		maxBodySize      int64
		wantsHealthCheck bool
		lock             sync.RWMutex
		total            uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
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

			m := &Mirroring{
				handler:          tt.fields.handler,
				mirrorHandlers:   tt.fields.mirrorHandlers,
				rw:               tt.fields.rw,
				routinePool:      tt.fields.routinePool,
				mirrorBody:       tt.fields.mirrorBody,
				maxBodySize:      tt.fields.maxBodySize,
				wantsHealthCheck: tt.fields.wantsHealthCheck,
				lock:             tt.fields.lock,
				total:            tt.fields.total,
			}
			if got := m.inc(); got != tt.want {
				t.Errorf("Mirroring.inc() = %v, want %v", got, tt.want)
			}
		})
	}
}
