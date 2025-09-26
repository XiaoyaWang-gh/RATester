package commands

import (
	"fmt"
	"testing"
)

func TestOnFileClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	detector := &fileChangeDetector{
		current: make(map[string]uint64),
		prev:    make(map[string]uint64),
	}

	detector.OnFileClose("test.txt", 123456)

	if detector.current["test.txt"] != 123456 {
		t.Errorf("Expected checksum 123456 for test.txt, but got %d", detector.current["test.txt"])
	}
}
