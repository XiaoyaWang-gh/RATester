package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbuildHealthCheckResponseList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	healthCheckResults := [][]string{
		{"service1", "OK", "200"},
		{"service2", "Error", "500"},
	}

	expectedResponse := []map[string]interface{}{
		{"name": "service1", "message": "OK", "status": "200"},
		{"name": "service2", "message": "Error", "status": "500"},
	}

	response := buildHealthCheckResponseList(&healthCheckResults)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected %v, but got %v", expectedResponse, response)
	}
}
