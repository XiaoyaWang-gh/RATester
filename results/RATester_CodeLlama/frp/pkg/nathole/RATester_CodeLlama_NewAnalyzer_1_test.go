package nathole

import (
	"fmt"
	"testing"
	"time"
)

func TestNewAnalyzer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := NewAnalyzer(time.Duration(10))
	if a == nil {
		t.Error("NewAnalyzer failed")
	}
}
