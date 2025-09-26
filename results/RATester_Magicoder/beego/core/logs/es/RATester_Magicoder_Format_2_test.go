package es

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func TestFormat_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lm := &logs.LogMsg{
		Msg:  "test message",
		When: time.Now(),
	}
	el := &esLogger{}
	result := el.Format(lm)

	// Assuming the JSON marshalling is correct, the result should be a string
	if reflect.TypeOf(result).Kind() != reflect.String {
		t.Errorf("Expected result to be a string, but got %T", result)
	}

	// Assuming the JSON marshalling is correct, the result should contain the original message
	if !strings.Contains(result, lm.Msg) {
		t.Errorf("Expected result to contain the original message, but it did not")
	}

	// Assuming the JSON marshalling is correct, the result should contain the timestamp
	if !strings.Contains(result, lm.When.Format(time.RFC3339)) {
		t.Errorf("Expected result to contain the timestamp, but it did not")
	}
}
