package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestContentWithoutSummary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	ctx := context.Background()
	_, err := pco.ContentWithoutSummary(ctx)
	if err != nil {
		t.Errorf("ContentWithoutSummary() error = %v", err)
		return
	}
}
