package goldmark

import (
	"fmt"
	"testing"
)

func TestnewIDFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	idType := "test"
	idFactory := newIDFactory(idType)

	if idFactory.idType != idType {
		t.Errorf("Expected idType to be %s, but got %s", idType, idFactory.idType)
	}

	if idFactory.vals == nil {
		t.Error("Expected vals to be initialized, but it was nil")
	}
}
