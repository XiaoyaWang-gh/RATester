package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewConnector_1(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg *v1.ClientCommonConfig
	}
	tests := []struct {
		name string
		args args
		want Connector
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

			if got := NewConnector(tt.args.ctx, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}
