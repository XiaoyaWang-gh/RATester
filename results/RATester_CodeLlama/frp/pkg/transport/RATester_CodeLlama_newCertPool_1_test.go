package transport

import (
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestNewCertPool_1(t *testing.T) {
	type args struct {
		caPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *x509.CertPool
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				caPath: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test case 2",
			args: args{
				caPath: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test case 3",
			args: args{
				caPath: "",
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

			got, err := newCertPool(tt.args.caPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCertPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCertPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
