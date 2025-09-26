package observability

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewStatusCodeRecorder_1(t *testing.T) {
	type args struct {
		rw     http.ResponseWriter
		status int
	}
	tests := []struct {
		name string
		args args
		want *statusCodeRecorder
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

			if got := newStatusCodeRecorder(tt.args.rw, tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStatusCodeRecorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
