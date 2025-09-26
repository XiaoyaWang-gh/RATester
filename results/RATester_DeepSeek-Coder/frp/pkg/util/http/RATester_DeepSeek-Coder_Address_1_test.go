package http

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestAddress_1(t *testing.T) {
	type fields struct {
		addr           string
		ln             net.Listener
		tlsCfg         *tls.Config
		router         *mux.Router
		hs             *http.Server
		authMiddleware mux.MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestAddress",
			fields: fields{
				addr: "localhost:8080",
			},
			want: "localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Server{
				addr:           tt.fields.addr,
				ln:             tt.fields.ln,
				tlsCfg:         tt.fields.tlsCfg,
				router:         tt.fields.router,
				hs:             tt.fields.hs,
				authMiddleware: tt.fields.authMiddleware,
			}
			if got := s.Address(); got != tt.want {
				t.Errorf("Server.Address() = %v, want %v", got, tt.want)
			}
		})
	}
}
