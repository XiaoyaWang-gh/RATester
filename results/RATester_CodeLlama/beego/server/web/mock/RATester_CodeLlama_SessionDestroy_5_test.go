package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	ctx := context.Background()
	sid := "123456"
	err := s.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy() error = %v, want nil", err)
		return
	}
}
