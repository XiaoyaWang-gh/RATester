package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestStatus_3(t *testing.T) {
	type fields struct {
		TCPMiddlewareInfo *runtime.TCPMiddlewareInfo
		Name              string
		Provider          string
		Type              string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				TCPMiddlewareInfo: &runtime.TCPMiddlewareInfo{
					Status: "test_status",
				},
				Name:     "test_name",
				Provider: "test_provider",
				Type:     "test_type",
			},
			want: "test_status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := tcpMiddlewareRepresentation{
				TCPMiddlewareInfo: tt.fields.TCPMiddlewareInfo,
				Name:              tt.fields.Name,
				Provider:          tt.fields.Provider,
				Type:              tt.fields.Type,
			}
			if got := m.status(); got != tt.want {
				t.Errorf("tcpMiddlewareRepresentation.status() = %v, want %v", got, tt.want)
			}
		})
	}
}
