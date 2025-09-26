package version

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetPrint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	wr := &bytes.Buffer{}

	// when
	err := GetPrint(wr)

	// then
	if err != nil {
		t.Errorf("GetPrint() error = %v", err)
		return
	}

	if wr.String() != versionTemplate {
		t.Errorf("GetPrint() = %v, want %v", wr.String(), versionTemplate)
	}
}
