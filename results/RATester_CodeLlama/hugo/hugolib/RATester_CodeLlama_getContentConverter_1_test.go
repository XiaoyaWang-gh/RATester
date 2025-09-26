package hugolib

import (
	"fmt"
	"testing"
)

func TestGetContentConverter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	pco := &pageContentOutput{}
	// When
	result, err := pco.getContentConverter()
	// Then
	if err != nil {
		t.Errorf("getContentConverter() error = %v", err)
		return
	}
	if result == nil {
		t.Errorf("getContentConverter() result = %v", result)
	}
}
