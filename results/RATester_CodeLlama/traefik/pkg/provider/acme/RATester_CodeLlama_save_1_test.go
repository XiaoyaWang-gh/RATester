package acme

import (
	"fmt"
	"testing"
)

func TestSave_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &LocalStore{}
	resolverName := "resolverName"
	storedData := &StoredData{}

	// when
	s.save(resolverName, storedData)

	// then
	// TODO
}
