package fiber

import (
	"fmt"
	"testing"
)

func TestGetRouteURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		route: &Route{
			Method: "GET",
			Path:   "/test",
		},
	}

	params := Map{"id": "123"}
	expectedURL := "/test?id=123"

	url, err := ctx.GetRouteURL("test", params)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if url != expectedURL {
		t.Errorf("Expected URL: %s, but got: %s", expectedURL, url)
	}
}
