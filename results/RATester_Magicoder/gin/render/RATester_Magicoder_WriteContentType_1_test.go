package render

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := Data{
		ContentType: "application/json",
		Data:        []byte("test data"),
	}

	w := httptest.NewRecorder()
	data.WriteContentType(w)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if string(body) != "application/json" {
		t.Errorf("Expected content type 'application/json', got '%s'", string(body))
	}
}
