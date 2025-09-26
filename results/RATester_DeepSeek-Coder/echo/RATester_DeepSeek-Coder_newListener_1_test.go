package echo

import (
	"fmt"
	"testing"
)

func TestNewListener_1(t *testing.T) {
	type args struct {
		address string
		network string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid network and address",
			args: args{
				address: "localhost:8080",
				network: "tcp",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid network",
			args: args{
				address: "localhost:8080",
				network: "invalid",
			},
			wantErr: true,
		},
		{
			name: "Test with invalid address",
			args: args{
				address: "invalid:8080",
				network: "tcp",
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

			_, err := newListener(tt.args.address, tt.args.network)
			if (err != nil) != tt.wantErr {
				t.Errorf("newListener() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
