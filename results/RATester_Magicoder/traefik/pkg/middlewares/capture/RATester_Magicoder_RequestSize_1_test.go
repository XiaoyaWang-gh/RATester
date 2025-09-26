package capture

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRequestSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	capture := &Capture{
		rr: &readCounter{
			source: ioutil.NopCloser(strings.NewReader("test")),
			size:   4,
		},
	}

	size := capture.RequestSize()
	if size != 4 {
		t.Errorf("Expected RequestSize to be 4, but got %d", size)
	}
}
