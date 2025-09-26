package conn

import (
	"fmt"
	"testing"
)

func TestNewSecret_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte{},
	}
	secret := NewSecret("password", conn)

	if secret.Password != "password" {
		t.Errorf("Expected password to be 'password', but got '%s'", secret.Password)
	}

	if secret.Conn != conn {
		t.Errorf("Expected conn to be %v, but got %v", conn, secret.Conn)
	}
}
