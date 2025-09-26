package tcp

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/challenge/tlsalpn01"
)

func TestAlpnV2_1(t *testing.T) {
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
			name: "empty protos",
			args: args{
				tree:   &matchersTree{},
				protos: []string{},
			},
			wantErr: true,
		},
		{
			name: "valid protos",
			args: args{
				tree:   &matchersTree{},
				protos: []string{"proto1", "proto2"},
			},
			wantErr: false,
		},
		{
			name: "invalid protocol",
			args: args{
				tree:   &matchersTree{},
				protos: []string{tlsalpn01.ACMETLS1Protocol},
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

			err := alpnV2(tt.args.tree, tt.args.protos...)
			if (err != nil) != tt.wantErr {
				t.Errorf("alpnV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
