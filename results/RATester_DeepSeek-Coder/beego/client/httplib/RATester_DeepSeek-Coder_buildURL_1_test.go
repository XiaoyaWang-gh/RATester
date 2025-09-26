package httplib

import (
	"fmt"
	"testing"
)

func TestBuildURL_1(t *testing.T) {
	type args struct {
		paramBody string
	}
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
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

			tt.b.buildURL(tt.args.paramBody)
		})
	}
}
