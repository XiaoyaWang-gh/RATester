package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	cookie1 := &http.Cookie{
		Name:  "test1",
		Value: "value1",
	}
	cookie2 := &http.Cookie{
		Name:  "test2",
		Value: "value2",
	}
	req.AddCookie(cookie1)
	req.AddCookie(cookie2)

	ctx := &context{
		request: req,
	}

	cookies := ctx.Cookies()

	if len(cookies) != 2 {
		t.Errorf("Expected 2 cookies, got %d", len(cookies))
	}

	if cookies[0].Name != "test1" || cookies[0].Value != "value1" {
		t.Errorf("Expected cookie test1 with value value1, got %s with value %s", cookies[0].Name, cookies[0].Value)
	}

	if cookies[1].Name != "test2" || cookies[1].Value != "value2" {
		t.Errorf("Expected cookie test2 with value value2, got %s with value %s", cookies[1].Name, cookies[1].Value)
	}
}
