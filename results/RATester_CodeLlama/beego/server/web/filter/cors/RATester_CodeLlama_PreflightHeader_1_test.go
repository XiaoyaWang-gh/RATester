package cors

import (
	"fmt"
	"testing"
	"time"
)

func TestPreflightHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Content-Type", "Accept", "Authorization"},
		ExposeHeaders:   []string{"Content-Length", "Content-Type"},
		MaxAge:          12 * time.Hour,
	}

	headers := o.PreflightHeader("http://example.com", "GET", "Content-Type, Accept")
	if len(headers) != 5 {
		t.Errorf("Expected 5 headers, got %d", len(headers))
	}

	if headers[headerAllowOrigin] != "*" {
		t.Errorf("Expected %s header to be *, got %s", headerAllowOrigin, headers[headerAllowOrigin])
	}

	if headers[headerAllowMethods] != "GET,POST,PUT,DELETE" {
		t.Errorf("Expected %s header to be GET,POST,PUT,DELETE, got %s", headerAllowMethods, headers[headerAllowMethods])
	}

	if headers[headerAllowHeaders] != "Content-Type,Accept" {
		t.Errorf("Expected %s header to be Content-Type,Accept, got %s", headerAllowHeaders, headers[headerAllowHeaders])
	}

	if headers[headerExposeHeaders] != "Content-Length,Content-Type" {
		t.Errorf("Expected %s header to be Content-Length,Content-Type, got %s", headerExposeHeaders, headers[headerExposeHeaders])
	}

	if headers[headerMaxAge] != "43200" {
		t.Errorf("Expected %s header to be 43200, got %s", headerMaxAge, headers[headerMaxAge])
	}
}
