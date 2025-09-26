package adaptor

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestHTTPHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// Act
	handler := HTTPHandlerFunc(h)
	// Assert
	assert.NotNil(t, handler)
}
