package page

import (
	"context"
	"fmt"
	"testing"
)

func TestContentWithoutSummary_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	nop := nopContent(0)
	_, err := nop.ContentWithoutSummary(ctx)
	if err != nil {
		t.Errorf("ContentWithoutSummary() error = %v, wantErr %v", err, false)
		return
	}
}
