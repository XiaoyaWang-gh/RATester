package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var trees methodTrees
	var method string
	var expected *node
	var actual *node
	// act
	actual = trees.get(method)
	// assert
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
