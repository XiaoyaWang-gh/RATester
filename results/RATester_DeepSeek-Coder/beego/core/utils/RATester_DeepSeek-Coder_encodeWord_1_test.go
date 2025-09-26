package utils

import (
	"fmt"
	"testing"
)

func TestEncodeWord_1(t *testing.T) {
	testCases := []struct {
		charset string
		s       string
		want    string
	}{
		{"UTF-8", "Hello, World", "=?UTF-8?q?Hello,_World?="},
		{"ISO-8859-1", "Hello, World", "=?ISO-8859-1?q?Hello,_World?="},
		{"US-ASCII", "Hello, World", "=?US-ASCII?q?Hello,_World?="},
		{"UTF-8", "Hello, =?UTF-8?q?World?=", "=?UTF-8?q?Hello,_=?UTF-8?q?World?=?="},
		{"UTF-8", "Hello, \x00", "=?UTF-8?q?Hello,_?="},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("charset=%s, s=%s", tc.charset, tc.s), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeWord(tc.charset, tc.s)
			if got != tc.want {
				t.Errorf("encodeWord(%q, %q) = %q, want %q", tc.charset, tc.s, got, tc.want)
			}
		})
	}
}
