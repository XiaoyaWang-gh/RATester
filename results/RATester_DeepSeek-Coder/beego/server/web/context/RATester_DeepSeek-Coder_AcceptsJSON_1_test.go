package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAcceptsJSON_1(t *testing.T) {
	type testCase struct {
		name   string
		header string
		want   bool
	}

	testCases := []testCase{
		{"case1", "application/json", true},
		{"case2", "application/xml", false},
		{"case3", "text/html", false},
		{"case4", "application/xhtml+xml", false},
		{"case5", "*/*", false},
		{"case6", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			input := &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Header: http.Header{
							"Accept": []string{tc.header},
						},
					},
				},
			}

			got := input.AcceptsJSON()
			if got != tc.want {
				t.Errorf("AcceptsJSON() = %v, want %v", got, tc.want)
			}
		})
	}
}
