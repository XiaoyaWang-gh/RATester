package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func Testclear_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &WhatChanged{
		identitySet: identity.Identities{},
	}

	w.clear()

	if len(w.identitySet) != 0 {
		t.Errorf("Expected identitySet to be empty, but it's not.")
	}
}
