package net

import (
	"fmt"
	"testing"
)

func TestListenKcp_1(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestListenKcp",
			args: args{
				address: "127.0.0.1:12345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := ListenKcp(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListenKcp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
