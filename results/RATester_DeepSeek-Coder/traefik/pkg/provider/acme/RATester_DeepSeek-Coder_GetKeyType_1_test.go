package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/certcrypto"
)

func TestGetKeyType_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		value string
	}
	tests := []struct {
		name string
		args args
		want certcrypto.KeyType
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

			if got := GetKeyType(tt.args.ctx, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeyType() = %v, want %v", got, tt.want)
			}
		})
	}
}
