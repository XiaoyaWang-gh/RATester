package legacy

import (
	"fmt"
	"testing"
)

func TestGetRenderedConfFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	path := "./testdata/test.conf"

	// when
	out, err := GetRenderedConfFromFile(path)

	// then
	if err != nil {
		t.Errorf("GetRenderedConfFromFile() error = %v", err)
		return
	}

	if string(out) != "test" {
		t.Errorf("GetRenderedConfFromFile() = %v, want %v", string(out), "test")
	}
}
