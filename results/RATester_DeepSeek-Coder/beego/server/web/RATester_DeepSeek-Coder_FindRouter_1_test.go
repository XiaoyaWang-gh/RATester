package web

import (
	"fmt"
	"reflect"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestFindRouter_1(t *testing.T) {
	type args struct {
		context *beecontext.Context
	}
	tests := []struct {
		name           string
		args           args
		wantRouterInfo *ControllerInfo
		wantIsFind     bool
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

			p := &ControllerRegister{}
			gotRouterInfo, gotIsFind := p.FindRouter(tt.args.context)
			if !reflect.DeepEqual(gotRouterInfo, tt.wantRouterInfo) {
				t.Errorf("FindRouter() gotRouterInfo = %v, want %v", gotRouterInfo, tt.wantRouterInfo)
			}
			if gotIsFind != tt.wantIsFind {
				t.Errorf("FindRouter() gotIsFind = %v, want %v", gotIsFind, tt.wantIsFind)
			}
		})
	}
}
