package authz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequirePermission_1(t *testing.T) {
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

			a := &BasicAuthorizer{}
			a.RequirePermission(tt.args.w)
		})
	}
}
