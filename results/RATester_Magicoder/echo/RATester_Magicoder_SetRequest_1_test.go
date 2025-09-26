package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetRequest_1(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		c    *context
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

			tt.c.SetRequest(tt.args.r)
		})
	}
}
