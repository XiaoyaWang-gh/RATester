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

	app := &App{}
	expectedStack := [][]*Route{}

	if !reflect.DeepEqual(app.Stack(), expectedStack) {
		t.Errorf("Expected stack to be %v, but got %v", expectedStack, app.Stack())
	}
}
