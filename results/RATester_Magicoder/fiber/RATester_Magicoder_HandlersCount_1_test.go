package fiber

import (
	"fmt"
	"testing"
)

func TestHandlersCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		handlersCount: 10,
	}

	count := app.HandlersCount()

	if count != 10 {
		t.Errorf("Expected handlers count to be 10, but got %d", count)
	}
}
