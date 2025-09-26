package bean

import (
	"fmt"
	"testing"
)

func TestNewTagAutoWireBeanFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	factory := NewTagAutoWireBeanFactory()
	if factory == nil {
		t.Errorf("Expected a non-nil factory, got nil")
	}

	if factory.Adapters == nil {
		t.Errorf("Expected a non-nil Adapters, got nil")
	}

	if factory.FieldTagParser == nil {
		t.Errorf("Expected a non-nil FieldTagParser, got nil")
	}
}
