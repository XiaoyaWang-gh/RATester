package logs

import (
	"fmt"
	"testing"
)

func TestGetWhenFormatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PatternLogFormatter{WhenFormat: "2006-01-02"}
	if p.getWhenFormatter() != "2006-01-02" {
		t.Errorf("getWhenFormatter() = %v, want %v", p.getWhenFormatter(), "2006-01-02")
	}
}
