package hugolib

import (
	"fmt"
	"testing"
)

func TestTestOptWithNFDOnDarwin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{}
	TestOptWithNFDOnDarwin()(config)

	if !config.NFDFormOnDarwin {
		t.Errorf("Expected NFDFormOnDarwin to be true, but it was false")
	}
}
