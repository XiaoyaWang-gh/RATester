package forwardedheaders

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRewrite_1(t *testing.T) {
	type args struct {
		outreq *http.Request
	}
	tests := []struct {
		name string
		x    *XForwarded
		args args
		want *http.Request
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

			tt.x.rewrite(tt.args.outreq)
			if !reflect.DeepEqual(tt.args.outreq, tt.want) {
				t.Errorf("rewrite() = %v, want %v", tt.args.outreq, tt.want)
			}
		})
	}
}
