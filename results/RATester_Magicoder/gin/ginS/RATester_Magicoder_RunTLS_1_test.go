package ginS

import (
	"fmt"
	"testing"
)

func TestRunTLS_1(t *testing.T) {
	type args struct {
		addr     string
		certFile string
		keyFile  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr:     "localhost:8080",
				certFile: "cert.pem",
				keyFile:  "key.pem",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr:     "invalid_address",
				certFile: "cert.pem",
				keyFile:  "key.pem",
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

			if err := RunTLS(tt.args.addr, tt.args.certFile, tt.args.keyFile); (err != nil) != tt.wantErr {
				t.Errorf("RunTLS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
