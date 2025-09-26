package captcha

import (
	"fmt"
	"net/http"
	"testing"
)

func TestVerifyReq_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		c    *Captcha
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

			if got := tt.c.VerifyReq(tt.args.req); got != tt.want {
				t.Errorf("Captcha.VerifyReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
