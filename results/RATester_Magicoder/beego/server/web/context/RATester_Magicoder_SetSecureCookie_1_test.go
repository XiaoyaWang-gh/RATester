package context

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

func TestSetSecureCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{
				Output: &BeegoOutput{},
			},
		},
	}

	Secret := "secret"
	name := "name"
	value := "value"

	ctx.SetSecureCookie(Secret, name, value)

	cookie := ctx.Input.Cookie(name)
	parts := strings.Split(cookie, "|")
	if len(parts) != 3 {
		t.Errorf("Expected 3 parts, got %d", len(parts))
	}

	vs, timestamp, sig := parts[0], parts[1], parts[2]
	h := hmac.New(sha256.New, []byte(Secret))
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)
	expectedSig := fmt.Sprintf("%02x", h.Sum(nil))

	if sig != expectedSig {
		t.Errorf("Expected signature %s, got %s", expectedSig, sig)
	}

	decodedValue, err := base64.URLEncoding.DecodeString(vs)
	if err != nil {
		t.Errorf("Failed to decode value: %v", err)
	}

	if string(decodedValue) != value {
		t.Errorf("Expected value %s, got %s", value, decodedValue)
	}
}
