package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWithCheckRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	redirect := func(req *http.Request, via []*http.Request) error {
		return nil
	}
	WithCheckRedirect(redirect)(client)

	if client.Setting.CheckRedirect == nil {
		t.Error("CheckRedirect is nil")
	}
}
