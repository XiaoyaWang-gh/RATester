package httpcache

import (
	"fmt"
	"testing"
	"time"
)

func TestMarshalJSON_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := PollConfig{
		For:     GlobMatcher{Excludes: []string{"exclude1", "exclude2"}},
		Disable: false,
		Low:     time.Second,
		High:    time.Minute,
	}

	expectedJSON := `{"For":{"Excludes":["exclude1","exclude2"]},"Disable":false,"Low":"1s","High":"1m"}`

	jsonBytes, err := config.MarshalJSON()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON: %s, but got: %s", expectedJSON, string(jsonBytes))
	}
}
