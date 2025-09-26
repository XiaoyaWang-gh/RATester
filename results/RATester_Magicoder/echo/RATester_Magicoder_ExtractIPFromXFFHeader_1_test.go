package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractIPFromXFFHeader_1(t *testing.T) {
	type args struct {
		options []TrustOption
		req     *http.Request
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

			if got := ExtractIPFromXFFHeader(tt.args.options...)(tt.args.req); got != tt.want {
				t.Errorf("ExtractIPFromXFFHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
