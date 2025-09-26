package hugolib

import (
	"errors"
	"fmt"
	"testing"
)

func TesterrorCollator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	results := make(chan error)
	errs := make(chan error)

	go s.errorCollator(results, errs)

	// Test case 1: No errors
	results <- nil
	close(results)
	err := <-errs
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// Test case 2: One error
	results = make(chan error)
	errs = make(chan error)
	go s.errorCollator(results, errs)
	results <- errors.New("Test error")
	close(results)
	err = <-errs
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 3: Multiple errors
	results = make(chan error)
	errs = make(chan error)
	go s.errorCollator(results, errs)
	results <- errors.New("Test error 1")
	results <- errors.New("Test error 2")
	close(results)
	err = <-errs
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
