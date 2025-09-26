package para

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := &errGroupRunner{
		Group: new(errgroup.Group),
		w:     &Workers{sem: make(chan struct{}, 1)},
		ctx:   ctx,
	}

	g.Run(func() error {
		return nil
	})

	g.Run(func() error {
		return errors.New("some error")
	})

	err := g.Wait()
	if err == nil {
		t.Fatal("expected error")
	}
}
