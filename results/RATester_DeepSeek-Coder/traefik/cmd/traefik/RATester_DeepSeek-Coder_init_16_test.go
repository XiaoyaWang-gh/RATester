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

	level := zerolog.GlobalLevel()

	if level != zerolog.ErrorLevel {
		t.Errorf("Expected level to be %v, got %v", zerolog.ErrorLevel, level)
	}
}
