package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

func TestRewriteURL_1(t *testing.T) {
	type args struct {
		rewriteRegex map[*regexp.Regexp]string
		req          *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			if err := rewriteURL(tt.args.rewriteRegex, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("rewriteURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
