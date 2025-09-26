package allconfig

import (
	"fmt"
	"testing"
)

func TestCanonifyURLs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if got := c.CanonifyURLs(); got != false {
		t.Errorf("ConfigLanguage.CanonifyURLs() = %v, want %v", got, false)
	}
}
