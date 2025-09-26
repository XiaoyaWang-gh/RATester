package retry

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewResponseWriter_1(t *testing.T) {
	type args struct {
		rw          http.ResponseWriter
		shouldRetry bool
	}
	tests := []struct {
		name string
		args args
		want *responseWriter
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

			got := newResponseWriter(tt.args.rw, tt.args.shouldRetry)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResponseWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
