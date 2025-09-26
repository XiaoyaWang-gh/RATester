package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ip := IP{
		Match: Match{
			Regexp: regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`),
		},
		Key: "ip",
	}

	expected := "IP"
	actual := ip.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
