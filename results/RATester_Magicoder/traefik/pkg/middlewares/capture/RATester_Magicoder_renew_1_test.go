package capture

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRenew_1(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name  string
		c     *Capture
		args  args
		want  http.ResponseWriter
		want1 *http.Request
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

			got, got1 := tt.c.renew(tt.args.rw, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renew() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("renew() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
