package grace

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListenAndServeTLS_2(t *testing.T) {
	type args struct {
		addr     string
		certFile string
		keyFile  string
		handler  http.Handler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr:     ":8080",
				certFile: "cert.pem",
				keyFile:  "key.pem",
				handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr:     ":8080",
				certFile: "invalid_cert.pem",
				keyFile:  "invalid_key.pem",
				handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := ListenAndServeTLS(tt.args.addr, tt.args.certFile, tt.args.keyFile, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("ListenAndServeTLS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
