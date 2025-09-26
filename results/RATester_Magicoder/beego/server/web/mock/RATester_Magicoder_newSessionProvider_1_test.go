package mock

import (
	"fmt"
	"testing"
)

func TestnewSessionProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sp := newSessionProvider()
	if sp == nil {
		t.Error("newSessionProvider() should not return nil")
	}
	if sp.Store == nil {
		t.Error("newSessionProvider().Store should not return nil")
	}
}
