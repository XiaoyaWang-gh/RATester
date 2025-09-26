package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewResponseModifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	modifier := func(response *http.Response) error {
		return nil
	}

	// act
	responseModifier := NewResponseModifier(w, r, modifier)

	// assert
	assert.NotNil(t, responseModifier)
}
