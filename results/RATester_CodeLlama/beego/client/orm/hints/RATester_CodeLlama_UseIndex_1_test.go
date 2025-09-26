package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var indexes []string
	hint := UseIndex(indexes...)
	if got := hint.GetKey(); got != KeyUseIndex {
		t.Errorf("UseIndex() = %v, want %v", got, KeyUseIndex)
	}
	if got := hint.GetValue(); !reflect.DeepEqual(got, indexes) {
		t.Errorf("UseIndex() = %v, want %v", got, indexes)
	}
}
