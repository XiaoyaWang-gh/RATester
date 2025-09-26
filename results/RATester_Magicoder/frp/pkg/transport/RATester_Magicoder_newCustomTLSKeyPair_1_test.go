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
			name: "Test case 1",
			args: args{
				certfile: "testdata/server.crt",
				keyfile:  "testdata/server.key",
			},
			want:    &tls.Certificate{},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				certfile: "testdata/non-existent.crt",
				keyfile:  "testdata/non-existent.key",
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
