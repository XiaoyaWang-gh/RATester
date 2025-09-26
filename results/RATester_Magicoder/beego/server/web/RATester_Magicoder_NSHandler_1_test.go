package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNSHandler_1(t *testing.T) {
	type args struct {
		rootpath string
		h        http.Handler
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSHandler(tt.args.rootpath, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
