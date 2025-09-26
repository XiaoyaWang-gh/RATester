package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetCheckRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}

	redirect := func(req *http.Request, via []*http.Request) error {
		return nil
	}

	b.SetCheckRedirect(redirect)

	if b.setting.CheckRedirect == nil {
		t.Errorf("Expected CheckRedirect to be set, but it was not")
	}
}
