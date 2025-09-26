package virtual

import (
	"context"
	"fmt"
	"testing"
)

func TestRun_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	c := &Client{}
	err := c.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
