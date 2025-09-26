package keyauth

import (
	"fmt"
	"testing"
)

func TestDefaultKeyLookup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var (
		keyLookup  = "source:name"
		authScheme = "Bearer"
	)
	extractor, err := DefaultKeyLookup(keyLookup, authScheme)
	if err != nil {
		t.Error(err)
	}
	if extractor == nil {
		t.Error("extractor is nil")
	}
}
