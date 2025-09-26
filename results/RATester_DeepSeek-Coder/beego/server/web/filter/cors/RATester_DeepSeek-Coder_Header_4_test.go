package cors

import (
	"fmt"
	"reflect"
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
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "X-Header"},
		MaxAge:           10 * time.Minute,
	}

	origin := "http://example.com"
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Methods":     "GET,POST",
		"Access-Control-Allow-Headers":     "Content-Type,Authorization",
		"Access-Control-Expose-Headers":    "Content-Length,X-Header",
		"Access-Control-Max-Age":           "600",
	}

	headers := o.Header(origin)

	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("Expected headers %v, but got %v", expectedHeaders, headers)
	}
}
