package transport

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestNewCustomTLSKeyPair_1(t *testing.T) {
	type args struct {
		certfile string
		keyfile  string
	}
	tests := []struct {
		name    string
		args    args
		want    *tls.Certificate
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				certfile: "testdata/cert.pem",
				keyfile:  "testdata/key.pem",
			},
			want:    &tls.Certificate{},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				certfile: "testdata/non_existent_cert.pem",
				keyfile:  "testdata/non_existent_key.pem",
			},
			want:    nil,
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

			got, err := newCustomTLSKeyPair(tt.args.certfile, tt.args.keyfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCustomTLSKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCustomTLSKeyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
