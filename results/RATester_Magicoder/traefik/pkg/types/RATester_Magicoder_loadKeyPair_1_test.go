package types

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestLoadKeyPair_1(t *testing.T) {
	type args struct {
		cert string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    tls.Certificate
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				cert: "testdata/server.crt",
				key:  "testdata/server.key",
			},
			want:    tls.Certificate{},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				cert: "testdata/server.crt",
				key:  "testdata/server.key",
			},
			want:    tls.Certificate{},
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

			got, err := loadKeyPair(tt.args.cert, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadKeyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
