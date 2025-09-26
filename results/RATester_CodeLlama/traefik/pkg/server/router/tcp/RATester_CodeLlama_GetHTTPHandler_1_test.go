package tcp

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetHTTPHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	r := &Router{}
	r.httpHandler = &http.ServeMux{}

	// act
	result := r.GetHTTPHandler()

	// assert
	assert.Equal(t, r.httpHandler, result)
}
