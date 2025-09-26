package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestPATCH_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.PATCH("/patch", func(c *gin.Context) {
		c.String(200, "PATCH method")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patch", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "PATCH method", w.Body.String())
}
