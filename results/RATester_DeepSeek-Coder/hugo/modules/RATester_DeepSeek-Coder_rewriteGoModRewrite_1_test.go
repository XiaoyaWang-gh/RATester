package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRewriteGoModRewrite_1(t *testing.T) {
	type args struct {
		name    string
		isGoMod map[string]bool
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    []byte
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

			got, err := tt.c.rewriteGoModRewrite(tt.args.name, tt.args.isGoMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.rewriteGoModRewrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.rewriteGoModRewrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
