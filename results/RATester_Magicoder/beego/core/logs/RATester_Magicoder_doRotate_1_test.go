package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestdoRotate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
		MaxFiles: 5,
	}

	err := w.doRotate(time.Now())
	if err != nil {
		t.Errorf("doRotate failed: %v", err)
	}

	// Add more test cases as needed
}
