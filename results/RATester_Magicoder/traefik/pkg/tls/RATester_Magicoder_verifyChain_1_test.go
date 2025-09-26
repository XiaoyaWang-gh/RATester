package tls

import (
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestVerifyChain_1(t *testing.T) {
	type args struct {
		rootCAs  *x509.CertPool
		rawCerts [][]byte
	}
	tests := []struct {
		name    string
		args    args
		want    *x509.Certificate
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				rootCAs:  nil,
				rawCerts: [][]byte{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Case 2",
			args: args{
				rootCAs:  nil,
				rawCerts: [][]byte{[]byte{}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				rootCAs:  nil,
				rawCerts: [][]byte{[]byte{0x00}},
			},
			want:    nil,
			wantErr: true,
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

			got, err := verifyChain(tt.args.rootCAs, tt.args.rawCerts)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verifyChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
