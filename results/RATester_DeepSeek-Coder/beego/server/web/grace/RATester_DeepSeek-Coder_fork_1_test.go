package grace

import (
	"fmt"
	"net"
	"net/http"
	"testing"
)

func TestFork_1(t *testing.T) {
	srv := &Server{
		Server: &http.Server{
			Addr: ":8080",
		},
		ln: &net.TCPListener{},
	}

	tests := []struct {
		name    string
		srv     *Server
		wantErr bool
	}{
		{
			name:    "Test fork",
			srv:     srv,
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

			err := tt.srv.fork()
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.fork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
