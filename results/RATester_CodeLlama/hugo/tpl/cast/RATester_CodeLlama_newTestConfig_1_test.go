package cast

import (
	"fmt"
	"testing"
)

func TestNewTestConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := newTestConfig()
	if v.Get("contentDir") != "content" {
		t.Errorf("Expected contentDir to be 'content', got %q", v.Get("contentDir"))
	}
}
