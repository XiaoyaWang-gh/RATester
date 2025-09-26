package service

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewInternalHandlers_1(t *testing.T) {
	type args struct {
		next           serviceManager
		apiHandler     http.Handler
		rest           http.Handler
		metricsHandler http.Handler
		pingHandler    http.Handler
		dashboard      http.Handler
		acmeHTTP       http.Handler
	}
	tests := []struct {
		name string
		args args
		want *InternalHandlers
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

			if got := NewInternalHandlers(tt.args.next, tt.args.apiHandler, tt.args.rest, tt.args.metricsHandler, tt.args.pingHandler, tt.args.dashboard, tt.args.acmeHTTP); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
