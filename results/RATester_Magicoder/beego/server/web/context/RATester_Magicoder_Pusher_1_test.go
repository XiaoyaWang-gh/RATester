package context

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestPusher_1(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		Started        bool
		Status         int
		Elapsed        time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Pusher
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
			if got := r.Pusher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Pusher() = %v, want %v", got, tt.want)
			}
		})
	}
}
