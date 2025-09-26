package auth

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSecretBasic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	b := &basicAuth{}
	user := "user"
	realm := "realm"

	// when
	result := b.secretBasic(user, realm)

	// then
	assert.Equal(t, "", result)
}
