package keyauth

import (
	"fmt"
	"testing"
)

func TestMultipleKeySourceLookup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	keyLookups := []string{"header", "query", "cookie"}
	authScheme := "Bearer"
	keyLookupFunc, err := MultipleKeySourceLookup(keyLookups, authScheme)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if keyLookupFunc == nil {
		t.Errorf("keyLookupFunc should not be nil")
	}
}
