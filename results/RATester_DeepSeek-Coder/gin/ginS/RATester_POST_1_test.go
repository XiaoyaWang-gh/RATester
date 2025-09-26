package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestPOST_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.POST("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
