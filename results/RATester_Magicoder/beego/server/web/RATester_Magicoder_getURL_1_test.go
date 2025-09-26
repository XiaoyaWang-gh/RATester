package web

import (
	"fmt"
	"testing"
)

func TestgetURL_1(t *testing.T) {
	type args struct {
		t              *Tree
		url            string
		controllerName string
		methodName     string
		params         map[string]string
		httpMethod     string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
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
			got, got1 := p.getURL(tt.args.t, tt.args.url, tt.args.controllerName, tt.args.methodName, tt.args.params, tt.args.httpMethod)
			if got != tt.want {
				t.Errorf("getURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
