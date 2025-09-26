package mirror

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestRegisterStatusUpdater_3(t *testing.T) {
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
	type args struct {
		fn func(up bool)
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
			if err := m.RegisterStatusUpdater(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Mirroring.RegisterStatusUpdater() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
