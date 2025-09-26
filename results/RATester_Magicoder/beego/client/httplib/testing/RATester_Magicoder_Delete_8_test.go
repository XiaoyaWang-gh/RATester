package testing

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestDelete_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "http://example.com/foo",
		httpmock.NewStringResponder(200, `{"id": 1, "name": "John Doe"}`))

	req := Delete("/foo")
	resp, err := req.Response()

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if result["id"] != 1.0 {
		t.Errorf("Expected id 1, got %f", result["id"])
	}

	if result["name"] != "John Doe" {
		t.Errorf("Expected name John Doe, got %s", result["name"])
	}
}
