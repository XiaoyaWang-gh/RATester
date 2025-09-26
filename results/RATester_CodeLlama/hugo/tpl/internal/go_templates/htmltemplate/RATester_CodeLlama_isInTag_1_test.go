package template

import (
	"fmt"
	"testing"
)

func TestIsInTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if isInTag(stateTag) != true {
		t.Errorf("isInTag(%v) = %v; want true", stateTag, isInTag(stateTag))
	}
	if isInTag(stateAttrName) != true {
		t.Errorf("isInTag(%v) = %v; want true", stateAttrName, isInTag(stateAttrName))
	}
	if isInTag(stateAfterName) != true {
		t.Errorf("isInTag(%v) = %v; want true", stateAfterName, isInTag(stateAfterName))
	}
	if isInTag(stateBeforeValue) != true {
		t.Errorf("isInTag(%v) = %v; want true", stateBeforeValue, isInTag(stateBeforeValue))
	}
	if isInTag(stateAttr) != true {
		t.Errorf("isInTag(%v) = %v; want true", stateAttr, isInTag(stateAttr))
	}
	if isInTag(state(0)) != false {
		t.Errorf("isInTag(%v) = %v; want false", state(0), isInTag(state(0)))
	}
}
