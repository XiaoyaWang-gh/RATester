package context

import (
	"fmt"
	"testing"
)

func TestCookie_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Context: &Context{
			ResponseWriter: &Response{},
		},
	}

	output.Cookie("name", "value", 100, "/path", "domain", true, true, "Lax")

	cookie := output.Context.ResponseWriter.Header().Get("Set-Cookie")

	if cookie != "name=value; Expires=; Max-Age=100; Path=/path; Domain=domain; Secure; HttpOnly; SameSite=Lax" {
		t.Errorf("Expected cookie to be 'name=value; Expires=; Max-Age=100; Path=/path; Domain=domain; Secure; HttpOnly; SameSite=Lax', but got '%s'", cookie)
	}
}
