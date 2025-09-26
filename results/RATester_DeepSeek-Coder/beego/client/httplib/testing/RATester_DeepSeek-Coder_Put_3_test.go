package testing

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPut_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "http://localhost:8080/path",
		httpmock.NewStringResponder(200, `{"status": "ok"}`))

	req := Put("/path")
	resp, err := req.Response()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}
}
