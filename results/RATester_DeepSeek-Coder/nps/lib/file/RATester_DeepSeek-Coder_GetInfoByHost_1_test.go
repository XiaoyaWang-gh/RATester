package file

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetInfoByHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.GetInfoByHost("example.com", req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = db.GetInfoByHost("non-existing-host", req)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
