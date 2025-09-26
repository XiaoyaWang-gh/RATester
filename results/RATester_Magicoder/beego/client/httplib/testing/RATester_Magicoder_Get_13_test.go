package testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGet_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/foo",
		httpmock.NewStringResponder(200, `{"foo": "bar"}`))

	resp, err := Get("http://example.com/foo").Response()
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result["foo"] != "bar" {
		t.Fatal("Expected foo to be bar")
	}
}
