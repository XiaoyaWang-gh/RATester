package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestResetColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	params := &LogFormatterParams{
		Request: &http.Request{},
	}

	expected := "\033[0m"
	actual := params.ResetColor()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
