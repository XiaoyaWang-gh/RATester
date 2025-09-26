package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestContainsHeader_1(t *testing.T) {
	type args struct {
		req   *http.Request
		name  string
		value string
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

			if got := containsHeader(tt.args.req, tt.args.name, tt.args.value); got != tt.want {
				t.Errorf("containsHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
