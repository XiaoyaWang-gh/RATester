package capture

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestReset_3(t *testing.T) {
	type fields struct {
		rr  *readCounter
		crw *captureResponseWriter
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
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

			c := &Capture{
				rr:  tt.fields.rr,
				crw: tt.fields.crw,
			}
			got := c.Reset(tt.args.next)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}
