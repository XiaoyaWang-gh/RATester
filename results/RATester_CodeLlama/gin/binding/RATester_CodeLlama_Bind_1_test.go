package binding

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `form:"name"`
	}

	req, _ := http.NewRequest("POST", "/", nil)
	req.Form = url.Values{}
	req.Form.Add("name", "test")

	var obj testStruct
	err := formMultipartBinding{}.Bind(req, &obj)
	if err != nil {
		t.Fatal(err)
	}

	if obj.Name != "test" {
		t.Fatalf("expected name to be test, got %s", obj.Name)
	}
}
