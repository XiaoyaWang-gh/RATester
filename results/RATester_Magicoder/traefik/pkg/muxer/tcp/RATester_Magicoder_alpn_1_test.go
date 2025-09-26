package tcp

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/challenge/tlsalpn01"
)

func TestAlpn_1(t *testing.T) {
	type args struct {
		tree   *matchersTree
		protos []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				tree:   &matchersTree{},
				protos: []string{tlsalpn01.ACMETLS1Protocol},
			},
			wantErr: true,
		},
		{
			name: "Test case 2",
			args: args{
				tree:   &matchersTree{},
				protos: []string{"valid_protocol"},
			},
			wantErr: false,
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

			if err := alpn(tt.args.tree, tt.args.protos...); (err != nil) != tt.wantErr {
				t.Errorf("alpn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
