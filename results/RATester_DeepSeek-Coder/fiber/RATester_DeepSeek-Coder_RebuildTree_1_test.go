package fiber

import (
	"fmt"
	"sync"
	"testing"
)

func TestRebuildTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := &App{
		mutex: sync.Mutex{},
	}

	expected := app.buildTree()
	actual := app.RebuildTree()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
