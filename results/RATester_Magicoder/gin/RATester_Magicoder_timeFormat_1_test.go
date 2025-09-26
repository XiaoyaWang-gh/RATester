package gin

import (
	"fmt"
	"testing"
	"time"
)

func TesttimeFormat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTime := time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC)
	expected := "2022/01/01 - 12:00:00"
	result := timeFormat(testTime)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
