package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetModlineSplitter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	isGoMod := true
	line := "require ( // indirect"
	expected := []string{"//", "indirect"}
	actual := getModlineSplitter(isGoMod)(line)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
