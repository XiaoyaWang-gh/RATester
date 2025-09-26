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
		t.Error("NewTagAutoWireBeanFactory() should not return nil")
	}

	if factory.Adapters == nil {
		t.Error("Adapters should not be nil")
	}

	if factory.FieldTagParser == nil {
		t.Error("FieldTagParser should not be nil")
	}
}
