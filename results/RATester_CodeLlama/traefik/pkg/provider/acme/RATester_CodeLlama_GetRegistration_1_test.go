package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/go-acme/lego/v4/registration"
)

func TestGetRegistration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	a := &Account{
		Registration: &registration.Resource{},
	}

	// when
	result := a.GetRegistration()

	// then
	assert.Equal(t, a.Registration, result)
}
