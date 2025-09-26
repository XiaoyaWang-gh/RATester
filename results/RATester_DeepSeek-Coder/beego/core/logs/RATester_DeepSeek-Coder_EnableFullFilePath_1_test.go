package logs

import (
	"fmt"
	"testing"
)

func TestEnableFullFilePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	EnableFullFilePath(true)
	if !beeLogger.enableFullFilePath {
		t.Errorf("Expected EnableFullFilePath to set enableFullFilePath to true")
	}

	EnableFullFilePath(false)
	if beeLogger.enableFullFilePath {
		t.Errorf("Expected EnableFullFilePath to set enableFullFilePath to false")
	}
}
