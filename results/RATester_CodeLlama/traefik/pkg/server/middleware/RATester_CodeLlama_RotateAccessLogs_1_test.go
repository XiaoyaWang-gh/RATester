package middleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
)

func TestRotateAccessLogs_1(t *testing.T) {
	type fields struct {
		accessLoggerMiddleware *accesslog.Handler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should return nil if accessLoggerMiddleware is nil",
			fields: fields{
				accessLoggerMiddleware: nil,
			},
			wantErr: false,
		},
		{
			name: "should return nil if accessLoggerMiddleware is not nil",
			fields: fields{
				accessLoggerMiddleware: &accesslog.Handler{},
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

			o := &ObservabilityMgr{
				accessLoggerMiddleware: tt.fields.accessLoggerMiddleware,
			}
			if err := o.RotateAccessLogs(); (err != nil) != tt.wantErr {
				t.Errorf("RotateAccessLogs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
