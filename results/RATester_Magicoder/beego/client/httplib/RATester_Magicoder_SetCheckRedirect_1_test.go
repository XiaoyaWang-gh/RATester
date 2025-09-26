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

	req := &BeegoHTTPRequest{
		setting: BeegoHTTPSettings{},
	}

	redirect := func(req *http.Request, via []*http.Request) error {
		return nil
	}

	req.SetCheckRedirect(redirect)

	if req.setting.CheckRedirect == nil {
		t.Error("CheckRedirect is not set")
	}
}
