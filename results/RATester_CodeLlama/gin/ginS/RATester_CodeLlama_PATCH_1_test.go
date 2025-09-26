package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPATCH_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	engine := gin.New()
	engine.PATCH("/test", func(c *gin.Context) {
		c.String(200, "test")
	})

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/test", nil)
	engine.ServeHTTP(w, req)

	// Assert
	if w.Code != 200 {
		t.Errorf("Expected response code 200 but got %d", w.Code)
	}
	if w.Body.String() != "test" {
		t.Errorf("Expected response body 'test' but got %s", w.Body.String())
	}
}
