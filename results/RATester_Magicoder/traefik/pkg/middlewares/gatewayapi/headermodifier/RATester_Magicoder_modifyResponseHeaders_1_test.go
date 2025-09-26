package headermodifier

import (
	"fmt"
	"net/http"
	"testing"
)

func TestmodifyResponseHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	modifier := &responseHeaderModifier{
		set: map[string]string{
			"Set-Header": "Set-Value",
		},
		add: map[string]string{
			"Add-Header": "Add-Value",
		},
		remove: []string{"Remove-Header"},
	}

	res := &http.Response{
		Header: make(http.Header),
	}

	err := modifier.modifyResponseHeaders(res)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if res.Header.Get("Set-Header") != "Set-Value" {
		t.Errorf("Expected Set-Header to be Set-Value, got %s", res.Header.Get("Set-Header"))
	}

	if res.Header.Get("Add-Header") != "Add-Value" {
		t.Errorf("Expected Add-Header to be Add-Value, got %s", res.Header.Get("Add-Header"))
	}

	if res.Header.Get("Remove-Header") != "" {
		t.Errorf("Expected Remove-Header to be empty, got %s", res.Header.Get("Remove-Header"))
	}
}
