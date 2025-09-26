package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestwhiteSpaceTrimmer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := whiteSpaceTrimmer()
	entry := &logg.Entry{
		Message: " test message ",
		Fields: []logg.Field{
			{Name: "field1", Value: " field1 value "},
			{Name: "field2", Value: "field2 value"},
		},
	}
	err := handler.HandleLog(entry)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if entry.Message != "test message" {
		t.Errorf("Expected message to be 'test message', but got '%s'", entry.Message)
	}
	for _, field := range entry.Fields {
		if s, ok := field.Value.(string); ok {
			if s != "field1 value" && s != "field2 value" {
				t.Errorf("Expected field value to be 'field1 value' or 'field2 value', but got '%s'", s)
			}
		} else {
			t.Errorf("Expected field value to be a string, but got %T", field.Value)
		}
	}
}
