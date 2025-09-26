package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		stack: [][]*Route{
			{
				{Method: "GET"},
				{Method: "POST"},
			},
			{
				{Method: "PUT"},
				{Method: "DELETE"},
			},
		},
	}

	expected := [][]*Route{
		{
			{Method: "GET"},
			{Method: "POST"},
		},
		{
			{Method: "PUT"},
			{Method: "DELETE"},
		},
	}

	result := app.Stack()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
