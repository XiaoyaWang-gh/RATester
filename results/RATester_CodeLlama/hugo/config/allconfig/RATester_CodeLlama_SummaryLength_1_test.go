package allconfig

import (
	"fmt"
	"testing"
)

func TestSummaryLength_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if got := c.SummaryLength(); got != 0 {
		t.Errorf("SummaryLength() = %v, want %v", got, 0)
	}
}
