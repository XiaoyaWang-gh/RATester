package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestUnSafeCopyOfStoredData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &LocalStore{}
	s.storedData = map[string]*StoredData{
		"resolverName": {
			Account: &Account{},
		},
	}

	// when
	result := s.unSafeCopyOfStoredData()

	// then
	assert.Equal(t, map[string]*StoredData{
		"resolverName": {
			Account: &Account{},
		},
	}, result)
}
