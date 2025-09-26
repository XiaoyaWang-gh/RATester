package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	type T struct {
		Name string `xml:"name"`
	}

	test := T{Name: "test"}

	err := c.XML(http.StatusOK, test)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	expected := `<T><name>test</name></T>`
	if rec.Body.String() != expected {
		t.Errorf("Expected response to match '%s', got '%s'", expected, rec.Body.String())
	}
}
