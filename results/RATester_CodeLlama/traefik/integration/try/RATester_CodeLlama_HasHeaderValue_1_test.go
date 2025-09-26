package try

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasHeaderValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := "Content-Type"
	value := "application/json"
	exactMatch := true

	res := &http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	err := HasHeaderValue(header, value, exactMatch)(res)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res = &http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	exactMatch = false
	err = HasHeaderValue(header, value, exactMatch)(res)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res = &http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	value = "application/xml"
	err = HasHeaderValue(header, value, exactMatch)(res)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	res = &http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	header = "Content-Length"
	err = HasHeaderValue(header, value, exactMatch)(res)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
