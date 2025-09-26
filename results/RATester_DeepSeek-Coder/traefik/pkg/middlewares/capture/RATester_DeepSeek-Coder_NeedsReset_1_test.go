package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNeedsReset_1(t *testing.T) {
	type fields struct {
		rr  *readCounter
		crw *captureResponseWriter
	}
	type args struct {
		rw http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			if got := c.NeedsReset(tt.args.rw); got != tt.want {
				t.Errorf("Capture.NeedsReset() = %v, want %v", got, tt.want)
			}
		})
	}
}
