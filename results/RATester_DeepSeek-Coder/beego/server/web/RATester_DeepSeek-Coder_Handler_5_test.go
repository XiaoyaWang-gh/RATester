package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_5(t *testing.T) {
	type args struct {
		rootpath string
		h        http.Handler
	}
	tests := []struct {
		name string
		n    *Namespace
		args args
		want *Namespace
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

			if got := tt.n.Handler(tt.args.rootpath, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
