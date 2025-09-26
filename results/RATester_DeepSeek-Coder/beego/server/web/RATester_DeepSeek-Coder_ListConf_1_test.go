package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListConf_1(t *testing.T) {
	tests := []struct {
		name string
		a    *adminController
		want http.ResponseWriter
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

			a := &adminController{
				Controller: tt.a.Controller,
				servers:    tt.a.servers,
			}
			a.ListConf()
			if got := a.Ctx.ResponseWriter; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
