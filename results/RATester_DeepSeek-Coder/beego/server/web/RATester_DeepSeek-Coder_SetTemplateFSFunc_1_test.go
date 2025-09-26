package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetTemplateFSFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testFunc := func() http.FileSystem {
		return nil
	}

	SetTemplateFSFunc(testFunc)

	if beeTemplateFS == nil {
		t.Errorf("Expected beeTemplateFS to be set, but it is nil")
	}
}
