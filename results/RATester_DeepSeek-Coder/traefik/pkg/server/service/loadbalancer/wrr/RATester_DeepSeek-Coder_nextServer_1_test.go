package wrr

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestNextServer_1(t *testing.T) {
	type fields struct {
		stickyCookie     *stickyCookie
		wantsHealthCheck bool
		handlersMu       sync.RWMutex
		handlerMap       map[string]*namedHandler
		handlers         []*namedHandler
		curDeadline      float64
		status           map[string]struct{}
		updaters         []func(bool)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *namedHandler
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

			b := &Balancer{
				stickyCookie:     tt.fields.stickyCookie,
				wantsHealthCheck: tt.fields.wantsHealthCheck,
				handlersMu:       tt.fields.handlersMu,
				handlerMap:       tt.fields.handlerMap,
				handlers:         tt.fields.handlers,
				curDeadline:      tt.fields.curDeadline,
				status:           tt.fields.status,
				updaters:         tt.fields.updaters,
			}
			got, err := b.nextServer()
			if (err != nil) != tt.wantErr {
				t.Errorf("Balancer.nextServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balancer.nextServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
