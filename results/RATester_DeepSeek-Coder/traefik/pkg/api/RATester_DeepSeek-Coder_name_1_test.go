package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestName_1(t *testing.T) {
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
				TCPMiddlewareInfo: &runtime.TCPMiddlewareInfo{},
				Name:              "test",
				Provider:          "test",
				Type:              "test",
			},
			want: "test",
		},
		// Add more test cases as needed
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
			if got := m.name(); got != tt.want {
				t.Errorf("tcpMiddlewareRepresentation.name() = %v, want %v", got, tt.want)
			}
		})
	}
}
