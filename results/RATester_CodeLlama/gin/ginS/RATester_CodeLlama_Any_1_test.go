package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	engine.ServeHTTP(w, req)
	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Hello World!" {
		t.Errorf("Expected body %s but got %s", "Hello World!", w.Body.String())
	}
}
