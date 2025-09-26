package modules

import (
	"fmt"
	"testing"
)

func TestKey_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mount{
		Source: "scss",
		Target: "assets/bootstrap/scss",
		Lang:   "en",
	}
	expected := "en/scss/assets/bootstrap/scss"
	if m.key() != expected {
		t.Errorf("expected %s, got %s", expected, m.key())
	}
}
