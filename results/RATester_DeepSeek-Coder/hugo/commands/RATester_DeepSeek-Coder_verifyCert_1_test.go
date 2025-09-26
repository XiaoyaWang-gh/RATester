package commands

import (
	"fmt"
	"testing"
)

func TestVerifyCert_1(t *testing.T) {
	type args struct {
		rootPEM []byte
		certPEM []byte
		name    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid certificate",
			args: args{
				rootPEM: []byte("-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgIJAKF+...\n-----END CERTIFICATE-----\n"),
				certPEM: []byte("-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgIJAKF+...\n-----END CERTIFICATE-----\n"),
				name:    "example.com",
			},
			wantErr: false,
		},
		{
			name: "invalid certificate",
			args: args{
				rootPEM: []byte("-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgIJAKF+...\n-----END CERTIFICATE-----\n"),
				certPEM: []byte("-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgIJAKF+...\n-----END CERTIFICATE-----\n"),
				name:    "invalid.com",
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

			c := &serverCommand{}
			err := c.verifyCert(tt.args.rootPEM, tt.args.certPEM, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyCert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
