package lazy

import (
	"context"
	"fmt"
	"testing"
)

func TestInitCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &Init{}
	ini.Add(func(ctx context.Context) (any, error) {
		return nil, nil
	})
	ini.Add(func(ctx context.Context) (any, error) {
		return nil, nil
	})
	ini.Add(func(ctx context.Context) (any, error) {
		return nil, nil
	})
	if got := ini.InitCount(); got != 3 {
		t.Errorf("InitCount() = %v, want %v", got, 3)
	}
}
