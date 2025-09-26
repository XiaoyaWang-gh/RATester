package visitor

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewVisitor_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		cfg       v1.VisitorConfigurer
		clientCfg *v1.ClientCommonConfig
		helper    Helper
	}
	tests := []struct {
		name        string
		args        args
		wantVisitor Visitor
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

			gotVisitor := NewVisitor(tt.args.ctx, tt.args.cfg, tt.args.clientCfg, tt.args.helper)
			if !reflect.DeepEqual(gotVisitor, tt.wantVisitor) {
				t.Errorf("NewVisitor() = %v, want %v", gotVisitor, tt.wantVisitor)
			}
		})
	}
}
