package compress

import (
	"fmt"
	"net/http"
	"testing"
)

func TestChooseHandler_1(t *testing.T) {
	type args struct {
		typ  string
		rw   http.ResponseWriter
		req  *http.Request
		comp *compress
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

			c := &compress{}
			c.chooseHandler(tt.args.typ, tt.args.rw, tt.args.req)
		})
	}
}
