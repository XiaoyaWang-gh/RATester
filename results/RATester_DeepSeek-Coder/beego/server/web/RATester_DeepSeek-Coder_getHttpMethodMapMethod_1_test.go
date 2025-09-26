package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetHttpMethodMapMethod_1(t *testing.T) {
	type args struct {
		httpMethod string
		ctMethod   string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
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
			if got := p.getHttpMethodMapMethod(tt.args.httpMethod, tt.args.ctMethod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHttpMethodMapMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
