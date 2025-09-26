package context

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestReset_9(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		Started        bool
		Status         int
		Elapsed        time.Duration
	}
	type args struct {
		rw http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			r := &Response{
				ResponseWriter: tt.fields.ResponseWriter,
				Started:        tt.fields.Started,
				Status:         tt.fields.Status,
				Elapsed:        tt.fields.Elapsed,
			}
			r.reset(tt.args.rw)
		})
	}
}
