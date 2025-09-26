package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestContent_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	ctx := context.Background()
	_, err := pco.Content(ctx)
	if err != nil {
		t.Errorf("Content() error = %v", err)
		return
	}
}
