package data

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

	m := http.Header{}
	m.Set("key", "value")
	if !hasHeaderValue(m, "key", "value") {
		t.Errorf("hasHeaderValue(m, \"key\", \"value\") = false, want true")
	}
	if hasHeaderValue(m, "key", "value2") {
		t.Errorf("hasHeaderValue(m, \"key\", \"value2\") = true, want false")
	}
	if hasHeaderValue(m, "key2", "value") {
		t.Errorf("hasHeaderValue(m, \"key2\", \"value\") = true, want false")
	}
}
