package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOverview_1(t *testing.T) {
	type args struct {
		rw      *httptest.ResponseRecorder
		request *http.Request
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

			h := &Handler{}
			h.getOverview(tt.args.rw, tt.args.request)
		})
	}
}
