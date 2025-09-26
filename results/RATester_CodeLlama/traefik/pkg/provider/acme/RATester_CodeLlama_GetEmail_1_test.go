package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetEmail_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Account{
		Email: "test@example.com",
	}
	assert.Equal(t, "test@example.com", a.GetEmail())
}
