package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestHEAD_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.HEAD("/test", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("HEAD", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
