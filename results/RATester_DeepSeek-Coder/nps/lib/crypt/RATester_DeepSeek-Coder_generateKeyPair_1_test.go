package crypt

import (
	"fmt"
	"testing"
)

func TestGenerateKeyPair_1(t *testing.T) {
	type args struct {
		CommonName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				CommonName: "test.com",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				CommonName: "",
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

			_, _, err := generateKeyPair(tt.args.CommonName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
