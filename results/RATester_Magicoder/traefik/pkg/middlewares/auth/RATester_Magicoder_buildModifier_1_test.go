package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBuildModifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fa := &forwardAuth{
		addAuthCookiesToResponse: map[string]struct{}{"cookie1": {}, "cookie2": {}},
	}

	authCookies := []*http.Cookie{
		{Name: "cookie1", Value: "value1"},
		{Name: "cookie2", Value: "value2"},
		{Name: "cookie3", Value: "value3"},
	}

	modifier := fa.buildModifier(authCookies)

	res := &http.Response{
		Header: http.Header{},
	}

	res.Header.Add("Set-Cookie", "cookie1=value1; Path=/; HttpOnly")
	res.Header.Add("Set-Cookie", "cookie2=value2; Path=/; HttpOnly")
	res.Header.Add("Set-Cookie", "cookie3=value3; Path=/; HttpOnly")

	err := modifier(res)
	if err != nil {
		t.Fatal(err)
	}

	expectedCookies := []string{
		"cookie1=value1; Path=/; HttpOnly",
		"cookie2=value2; Path=/; HttpOnly",
		"cookie3=value3; Path=/; HttpOnly",
	}

	for _, expectedCookie := range expectedCookies {
		if !strings.Contains(res.Header.Get("Set-Cookie"), expectedCookie) {
			t.Errorf("Expected cookie %q not found in Set-Cookie header", expectedCookie)
		}
	}
}
