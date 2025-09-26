package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComposeControllerInfos_1(t *testing.T) {
	tests := []struct {
		name            string
		tree            *Tree
		wantRouterInfos []*ControllerInfo
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

			gotRouterInfos := []*ControllerInfo{}
			composeControllerInfos(tt.tree, &gotRouterInfos)
			if !reflect.DeepEqual(gotRouterInfos, tt.wantRouterInfos) {
				t.Errorf("composeControllerInfos() = %v, want %v", gotRouterInfos, tt.wantRouterInfos)
			}
		})
	}
}
