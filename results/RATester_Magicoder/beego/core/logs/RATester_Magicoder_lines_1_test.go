package logs

import (
	"fmt"
	"testing"
)

func Testlines_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}

	count, err := w.lines()
	if err != nil {
		t.Errorf("Error in lines: %v", err)
	}

	if count < 0 {
		t.Errorf("Invalid line count: %d", count)
	}
}
