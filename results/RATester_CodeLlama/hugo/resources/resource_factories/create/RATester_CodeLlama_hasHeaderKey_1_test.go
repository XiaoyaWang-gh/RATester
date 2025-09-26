package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasHeaderKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := http.Header{}
	m.Set("key", "value")
	if !hasHeaderKey(m, "key") {
		t.Errorf("hasHeaderKey(m, \"key\") = false, want true")
	}
	if hasHeaderKey(m, "nokey") {
		t.Errorf("hasHeaderKey(m, \"nokey\") = true, want false")
	}
}
