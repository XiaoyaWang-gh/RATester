package data

import (
	"fmt"
	"testing"
)

func TestParseCSV_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := []byte("a,b,c\nd,e,f\ng,h,i")
	sep := "\n"
	if _, err := parseCSV(c, sep); err != nil {
		t.Errorf("parseCSV(%q, %q) = %v", c, sep, err)
	}
}
