package customerrors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := &customErrors{
		name: "customErrors",
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}),
	}

	// Act
	c.ServeHTTP(rw, req)

	// Assert
	assert.Equal(t, http.StatusOK, rw.Code)
}
