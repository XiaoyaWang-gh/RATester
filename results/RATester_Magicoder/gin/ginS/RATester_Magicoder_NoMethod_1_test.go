package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestNoMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, "Method Not Allowed", w.Body.String())
}
