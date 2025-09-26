package langs

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{
		Lang: "en",
	}

	if got := l.String(); got != "en" {
		t.Errorf("String() = %v, want %v", got, "en")
	}
}
