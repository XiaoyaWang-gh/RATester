package observability

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFlush_3(t *testing.T) {
	type args struct {
		w http.ResponseWriter
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

			s := &statusCodeRecorder{
				ResponseWriter: tt.args.w,
			}
			s.Flush()
		})
	}
}
