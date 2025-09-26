package ginS

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStaticFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	relativePath := "/test"
	filepath := "../test.txt"

	r := gin.Default()
	r.StaticFile(relativePath, filepath)

	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected to get status %d but instead got %d", http.StatusOK, w.Code)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type to be %s but instead got %s", expectedContentType, w.Header().Get("Content-Type"))
	}

	expectedBody, _ := ioutil.ReadFile(filepath)
	if w.Body.String() != string(expectedBody) {
		t.Errorf("Expected body to be %s but instead got %s", string(expectedBody), w.Body.String())
	}
}
