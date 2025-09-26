package ginS

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStaticFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.StaticFS("/", http.Dir("./testdata"))

	srv := httptest.NewServer(router)
	defer srv.Close()

	res, err := http.Get(srv.URL + "/test.txt")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expectedBody := "This is a test file."
	if string(body) != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, string(body))
	}
}
