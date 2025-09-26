package capture

import (
	"fmt"
	"testing"
)

func TestHeader_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{}
	crw.Header()
	if crw.rw == nil {
		t.Errorf("crw.rw is nil")
	}
}
