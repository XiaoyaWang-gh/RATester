package web

import (
	"fmt"
	"testing"
)

func TestBuildHealthCheckResponseList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	healthCheckResults := [][]string{
		{"name1", "message1", "status1"},
		{"name2", "message2", "status2"},
		{"name3", "message3", "status3"},
	}

	response := buildHealthCheckResponseList(&healthCheckResults)

	if len(response) != 3 {
		t.Errorf("Expected response length to be 3, but got %d", len(response))
	}

	if response[0]["name"] != "name1" {
		t.Errorf("Expected response[0][\"name\"] to be \"name1\", but got %s", response[0]["name"])
	}

	if response[1]["message"] != "message2" {
		t.Errorf("Expected response[1][\"message\"] to be \"message2\", but got %s", response[1]["message"])
	}

	if response[2]["status"] != "status3" {
		t.Errorf("Expected response[2][\"status\"] to be \"status3\", but got %s", response[2]["status"])
	}
}
