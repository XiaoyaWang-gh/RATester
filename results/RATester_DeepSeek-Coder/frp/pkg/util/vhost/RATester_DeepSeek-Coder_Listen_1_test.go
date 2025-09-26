package vhost

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestListen_1(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg *RouteConfig
	}
	tests := []struct {
		name    string
		args    args
		wantL   *Listener
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

			v := &Muxer{}
			gotL, err := v.Listen(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Muxer.Listen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("Muxer.Listen() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
