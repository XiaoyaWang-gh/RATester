package acme

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveAccount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &LocalStore{}
	resolverName := "resolverName"
	account := &Account{}

	// when
	err := s.SaveAccount(resolverName, account)

	// then
	require.NoError(t, err)
}
