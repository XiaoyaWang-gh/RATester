package gin

import (
	"fmt"
	"testing"
)

func TestAuthorizationHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	user := "user"
	password := "password"
	want := "Basic dXNlcjpwYXNzd29yZA=="
	got := authorizationHeader(user, password)
	if got != want {
		t.Errorf("authorizationHeader(%s, %s) = %s; want %s", user, password, got, want)
	}
}
