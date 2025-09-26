package client

import (
	"fmt"
	"testing"
)

func TestAddMissingPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		addr   string
		isTLS  bool
		expect string
	}

	tests := []test{
		{"localhost", false, "localhost:80"},
		{"localhost", true, "localhost:443"},
		{"localhost:8080", false, "localhost:8080"},
		{"localhost:8080", true, "localhost:8080"},
		{"[::1]", false, "[::1]:80"},
		{"[::1]", true, "[::1]:443"},
		{"[::1]:8080", false, "[::1]:8080"},
		{"[::1]:8080", true, "[::1]:8080"},
	}

	for _, test := range tests {
		result := addMissingPort(test.addr, test.isTLS)
		if result != test.expect {
			t.Errorf("addMissingPort(%s, %t) = %s; want %s", test.addr, test.isTLS, result, test.expect)
		}
	}
}
