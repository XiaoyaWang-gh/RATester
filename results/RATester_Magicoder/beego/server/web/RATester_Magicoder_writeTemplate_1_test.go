package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestwriteTemplate_1(t *testing.T) {
	type args struct {
		rw   http.ResponseWriter
		data map[interface{}]interface{}
		tpls []string
	}
	tests := []struct {
		name string
		args args
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

			writeTemplate(tt.args.rw, tt.args.data, tt.args.tpls...)
		})
	}
}
