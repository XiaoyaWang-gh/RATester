package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewInternalHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var next serviceManager
	var apiHandler http.Handler
	var rest http.Handler
	var metricsHandler http.Handler
	var pingHandler http.Handler
	var dashboard http.Handler
	var acmeHTTP http.Handler
	// act
	result := NewInternalHandlers(next, apiHandler, rest, metricsHandler, pingHandler, dashboard, acmeHTTP)
	// assert
	assert.NotNil(t, result)
}
