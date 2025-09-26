package auth

import (
	"fmt"
	"testing"
)

func TestSecretDigest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &digestAuth{}
	d.users = make(map[string]string)
	d.users["user:realm"] = "secret"
	if got := d.secretDigest("user", "realm"); got != "secret" {
		t.Errorf("secretDigest() = %v, want %v", got, "secret")
	}
}
