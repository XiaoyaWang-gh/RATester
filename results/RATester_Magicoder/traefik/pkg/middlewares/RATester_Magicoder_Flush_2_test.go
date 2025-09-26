package middlewares

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFlush_2(t *testing.T) {
	type fields struct {
		req         *http.Request
		rw          http.ResponseWriter
		headersSent bool
		code        int
		modifier    func(*http.Response) error
		modified    bool
		modifierErr error
	}
	tests := []struct {
		name   string
		fields fields
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

			r := &ResponseModifier{
				req:         tt.fields.req,
				rw:          tt.fields.rw,
				headersSent: tt.fields.headersSent,
				code:        tt.fields.code,
				modifier:    tt.fields.modifier,
				modified:    tt.fields.modified,
				modifierErr: tt.fields.modifierErr,
			}
			r.Flush()
		})
	}
}
