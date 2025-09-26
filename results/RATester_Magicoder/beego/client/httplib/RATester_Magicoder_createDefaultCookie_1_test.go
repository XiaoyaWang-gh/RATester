package httplib

import (
	"fmt"
	"testing"
)

func TestcreateDefaultCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	createDefaultCookie()

	// Assert that defaultCookieJar is not nil
	if defaultCookieJar == nil {
		t.Errorf("Expected defaultCookieJar to be not nil")
	}
}
