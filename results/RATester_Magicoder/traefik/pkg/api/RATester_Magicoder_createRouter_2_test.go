package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestcreateRouter_2(t *testing.T) {
	type args struct {
		h Handler
	}
	tests := []struct {
		name string
		args args
		want *mux.Router
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

			if got := tt.args.h.createRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
