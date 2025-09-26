package fiber

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-Host", "test.com,example.com")
	resp, _ := app.Test(req)
	assert.Equal(t, "test.com", resp.Request.Host)
}
