package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetMethod_1(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getMethod(tt.args.r); got != tt.want {
				t.Errorf("getMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
