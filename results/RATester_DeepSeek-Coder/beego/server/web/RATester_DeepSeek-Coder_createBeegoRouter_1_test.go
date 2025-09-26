package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateBeegoRouter_1(t *testing.T) {
	type args struct {
		ct      reflect.Type
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *ControllerInfo
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
			if got := p.createBeegoRouter(tt.args.ct, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBeegoRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
