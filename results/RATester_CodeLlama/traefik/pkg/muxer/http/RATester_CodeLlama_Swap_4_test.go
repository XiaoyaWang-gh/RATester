package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSwap_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := routes{
		&route{
			matchers: matchersTree{
				matcher: func(req *http.Request) bool {
					return true
				},
			},
		},
		&route{
			matchers: matchersTree{
				matcher: func(req *http.Request) bool {
					return true
				},
			},
		},
	}
	r.Swap(0, 1)
	if r[0].matchers.matcher != nil {
		t.Errorf("Swap failed")
	}
}
