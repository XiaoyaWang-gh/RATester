package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	ctx := context.Background()
	s.SessionGC(ctx)
}
