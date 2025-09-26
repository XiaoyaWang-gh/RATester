package common

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckAuth_1(t *testing.T) {
	type args struct {
		r      *http.Request
		user   string
		passwd string
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

			if got := CheckAuth(tt.args.r, tt.args.user, tt.args.passwd); got != tt.want {
				t.Errorf("CheckAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
