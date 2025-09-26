package warpc

import (
	"context"
	"fmt"
	"testing"
)

func TestStart_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	type Q string
	type R string

	opts := Options{
		Ctx: context.Background(),
		Infof: func(format string, v ...any) {
			t.Logf(format, v...)
		},
		Runtime: Binary{
			Name: "testRuntime",
			Data: []byte("testData"),
		},
		Main: Binary{
			Name: "testMain",
			Data: []byte("testData"),
		},
		CompilationCacheDir: "testCacheDir",
		PoolSize:            10,
		Memory:              128,
	}

	d := &lazyDispatcher[Q, R]{
		opts: opts,
	}

	_, err := d.start()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test that start() is idempotent
	_, err = d.start()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
