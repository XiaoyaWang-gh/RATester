package authz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckPermission_1(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want bool
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
			if got := a.CheckPermission(tt.args.r); got != tt.want {
				t.Errorf("BasicAuthorizer.CheckPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}
