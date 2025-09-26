package tcp

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/challenge/tlsalpn01"
)

func TestAlpnV2_1(t *testing.T) {
	type test struct {
		name    string
		protos  []string
		wantErr bool
	}

	tests := []test{
		{
			name:    "empty protos",
			protos:  []string{},
			wantErr: true,
		},
		{
			name:    "invalid protocol",
			protos:  []string{tlsalpn01.ACMETLS1Protocol},
			wantErr: true,
		},
		{
			name:    "valid protocol",
			protos:  []string{"valid-protocol"},
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

			tree := &matchersTree{}
			err := alpnV2(tree, tt.protos...)
			if (err != nil) != tt.wantErr {
				t.Errorf("alpnV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
