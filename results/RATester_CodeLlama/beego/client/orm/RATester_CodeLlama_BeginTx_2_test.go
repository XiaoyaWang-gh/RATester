package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestBeginTx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	d := &DB{}
	tx, err := d.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("tx is nil")
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
