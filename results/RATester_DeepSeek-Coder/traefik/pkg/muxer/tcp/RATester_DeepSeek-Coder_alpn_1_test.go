package tcp

import (
	"fmt"
	"testing"
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
		// TODO: Add test cases.
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
