package bean

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDefaultValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	adapter := &TimeTypeAdapter{Layout: "2006-01-02"}

	// Test case 1: dftValue is "now"
	now := time.Now()
	value, err := adapter.DefaultValue(ctx, "now")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	actual, ok := value.(time.Time)
	if !ok {
		t.Errorf("Expected time.Time, got %T", value)
	}
	if actual.Before(now) || actual.After(now.Add(time.Second)) {
		t.Errorf("Expected time to be within a second of now, got %v", actual)
	}

	// Test case 2: dftValue is a valid time string
	dftValue := "2022-01-01"
	expected, err := time.Parse(adapter.Layout, dftValue)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	value, err = adapter.DefaultValue(ctx, dftValue)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	actual, ok = value.(time.Time)
	if !ok {
		t.Errorf("Expected time.Time, got %T", value)
	}
	if !actual.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test case 3: dftValue is an invalid time string
	dftValue = "invalid"
	_, err = adapter.DefaultValue(ctx, dftValue)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
