package mock

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMockByPath_1(t *testing.T) {
	type args struct {
		path string
		resp *http.Response
		err  error
	}
	tests := []struct {
		name string
		m    *MockResponseFilter
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

			tt.m.MockByPath(tt.args.path, tt.args.resp, tt.args.err)
		})
	}
}
