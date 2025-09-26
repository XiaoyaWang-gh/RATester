package failover

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetHandler_1(t *testing.T) {
	type args struct {
		handler http.Handler
	}
	tests := []struct {
		name string
		f    *Failover
		args args
		want http.Handler
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

			tt.f.SetHandler(tt.args.handler)
		})
	}
}
