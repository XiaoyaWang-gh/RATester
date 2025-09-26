package virtual

import (
	"fmt"
	"testing"
)

func TestOpen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := &pipeConnector{}
	err := pc.Open()
	if err != nil {
		t.Errorf("Open() failed: %v", err)
	}
}
