package sheap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &IntHeap{1, 2, 3}
	h.Push(4)
	want := IntHeap{1, 2, 3, 4}
	if !reflect.DeepEqual(*h, want) {
		t.Errorf("h.Push(4) = %v, want %v", *h, want)
	}
}
