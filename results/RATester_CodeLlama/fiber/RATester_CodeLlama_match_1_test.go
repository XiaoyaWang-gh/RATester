package fiber

import (
	"fmt"
	"testing"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// ### setup ###
	r := &Route{
		path: "/",
		use:  true,
	}
	detectionPath := "/"
	path := "/"
	params := &[maxParams]string{}
	// ### test ###
	if r.match(detectionPath, path, params) != true {
		t.Errorf("Expected true, got %v", r.match(detectionPath, path, params))
	}
	// ### verify ###
}
