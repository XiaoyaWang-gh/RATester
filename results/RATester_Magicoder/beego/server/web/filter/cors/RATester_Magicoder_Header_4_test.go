package cors

import (
	"fmt"
	"testing"
	"time"
)

func TestHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600 * time.Second,
	}

	headers := o.Header("http://example.com")

	if headers[headerAllowOrigin] != "*" {
		t.Errorf("Expected %s, got %s", "*", headers[headerAllowOrigin])
	}

	if headers[headerAllowCredentials] != "true" {
		t.Errorf("Expected %s, got %s", "true", headers[headerAllowCredentials])
	}

	if headers[headerAllowMethods] != "GET,POST" {
		t.Errorf("Expected %s, got %s", "GET,POST", headers[headerAllowMethods])
	}

	if headers[headerAllowHeaders] != "Content-Type,Authorization" {
		t.Errorf("Expected %s, got %s", "Content-Type,Authorization", headers[headerAllowHeaders])
	}

	if headers[headerExposeHeaders] != "Content-Length,Content-Type" {
		t.Errorf("Expected %s, got %s", "Content-Length,Content-Type", headers[headerExposeHeaders])
	}

	if headers[headerMaxAge] != "3600" {
		t.Errorf("Expected %s, got %s", "3600", headers[headerMaxAge])
	}
}
