package gin

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
)

func TestBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	accounts := Accounts{
		"admin": "password",
		"user":  "pass",
	}

	handler := BasicAuth(accounts)

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := &Context{Request: req}

	handler(ctx)

	if ctx.Writer.Status() != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, ctx.Writer.Status())
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:password"))
	req.Header.Set("Authorization", auth)

	handler(ctx)

	if ctx.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, ctx.Writer.Status())
	}
}
