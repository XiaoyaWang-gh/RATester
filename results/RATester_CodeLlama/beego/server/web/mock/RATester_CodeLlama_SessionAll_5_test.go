package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	ctx := context.Background()
	if s.SessionAll(ctx) != 0 {
		t.Errorf("SessionAll() = %v, want %v", s.SessionAll(ctx), 0)
	}
}
