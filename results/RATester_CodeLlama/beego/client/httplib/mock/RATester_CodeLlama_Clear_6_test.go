package mock

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClear_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MockResponseFilter{}
	m.MockByPath("/path", &http.Response{}, nil)
	m.Clear()
	if len(m.ms) != 0 {
		t.Errorf("m.ms should be empty")
	}
}
