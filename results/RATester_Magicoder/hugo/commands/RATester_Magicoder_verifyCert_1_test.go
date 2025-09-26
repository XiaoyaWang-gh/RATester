package commands

import (
	"fmt"
	"testing"
)

func TestverifyCert_1(t *testing.T) {
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
			name: "Test Case 1",
			args: args{
				rootPEM: []byte("-----BEGIN CERTIFICATE-----"),
				certPEM: []byte("-----BEGIN CERTIFICATE-----"),
				name:    "example.com",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				rootPEM: []byte("-----BEGIN CERTIFICATE-----"),
				certPEM: []byte("-----BEGIN CERTIFICATE-----"),
				name:    "example.com",
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
			if err := c.verifyCert(tt.args.rootPEM, tt.args.certPEM, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("verifyCert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
