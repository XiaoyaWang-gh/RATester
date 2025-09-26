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

	// Arrange
	// Act
	createDefaultCookie()
	// Assert
	if defaultCookieJar == nil {
		t.Error("defaultCookieJar is nil")
	}
}
