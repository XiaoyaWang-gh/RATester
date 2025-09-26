package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestTruncated_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pco := &pageContentOutput{
		po: &pageOutput{
			p: &pageState{
				pid: 1,
			},
		},
	}
	expected := true
	actual := pco.Truncated(ctx)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
