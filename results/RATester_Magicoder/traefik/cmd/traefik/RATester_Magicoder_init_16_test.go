package main

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestInit_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	if zerolog.GlobalLevel() != zerolog.ErrorLevel {
		t.Errorf("Expected zerolog.GlobalLevel() to be zerolog.ErrorLevel, but got %v", zerolog.GlobalLevel())
	}
}
