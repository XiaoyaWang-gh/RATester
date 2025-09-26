package hugolib

import (
	"fmt"
	"testing"
)

func TestAddShortcode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &contentParseInfo{}
	s := &shortcode{}
	p.AddShortcode(s)
	if len(p.itemsStep2) != 1 {
		t.Errorf("Expected 1 item in itemsStep2, got %d", len(p.itemsStep2))
	}
	if p.itemsStep2[0] != s {
		t.Errorf("Expected %v in itemsStep2, got %v", s, p.itemsStep2[0])
	}
}
