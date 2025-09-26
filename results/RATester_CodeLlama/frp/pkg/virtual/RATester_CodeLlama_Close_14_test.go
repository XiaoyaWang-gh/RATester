package virtual

import (
	"fmt"
	"testing"
)

func TestClose_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := &pipeConnector{}
	pc.Close()
}
