package mock

import (
	"fmt"
	"testing"
)

func TestWithJsonBodyFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := "field"
	value := "value"
	sc := &SimpleCondition{}
	WithJsonBodyFields(field, value)(sc)
	if sc.body[field] != value {
		t.Errorf("WithJsonBodyFields failed")
	}
}
