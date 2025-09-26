package httplib

import (
	"fmt"
	"testing"
)

func TestCreateDefaultCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	createDefaultCookie()

	if defaultCookieJar == nil {
		t.Errorf("Expected defaultCookieJar to be initialized, but it is nil")
	}
}
