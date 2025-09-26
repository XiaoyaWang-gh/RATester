package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestData_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		sd: ResourceSourceDescriptor{
			Data: map[string]any{
				"key": "value",
			},
		},
	}

	expected := map[string]any{
		"key": "value",
	}

	result := l.Data()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
