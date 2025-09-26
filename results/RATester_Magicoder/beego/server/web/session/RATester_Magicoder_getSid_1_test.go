package session

import (
	"fmt"
	"net/http"
	"testing"
)

func TestgetSid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		config: &ManagerConfig{
			EnableSidInURLQuery: true,
		},
	}

	req, err := http.NewRequest("GET", "/test?sid=testSid", nil)
	if err != nil {
		t.Fatal(err)
	}

	sid, err := manager.getSid(req)
	if err != nil {
		t.Fatal(err)
	}

	if sid != "testSid" {
		t.Fatalf("Expected sid to be 'testSid', but got '%s'", sid)
	}
}
