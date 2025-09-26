package utils

import (
	"fmt"
	"testing"
)

func TestQEncode_1(t *testing.T) {
	testCases := []struct {
		charset string
		s       string
		want    string
	}{
		{"utf-8", "hello", "hello"},
		{"utf-8", "world", "world"},
		{"utf-8", "你好", "%E4%BD%A0%E5%93%88"},
		{"utf-8", "世界", "%E4%B8%96%E7%95%8C"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("qEncode(%s, %s)", tc.charset, tc.s), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := qEncode(tc.charset, tc.s)
			if got != tc.want {
				t.Errorf("qEncode(%s, %s) = %s, want %s", tc.charset, tc.s, got, tc.want)
			}
		})
	}
}
