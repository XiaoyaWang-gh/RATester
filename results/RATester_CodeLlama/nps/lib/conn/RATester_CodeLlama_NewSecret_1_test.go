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

	// Arrange
	p := "password"
	conn := &Conn{}

	// Act
	secret := NewSecret(p, conn)

	// Assert
	if secret.Password != p {
		t.Errorf("Expected %s, got %s", p, secret.Password)
	}
	if secret.Conn != conn {
		t.Errorf("Expected %v, got %v", conn, secret.Conn)
	}
}
