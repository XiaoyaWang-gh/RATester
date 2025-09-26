package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	ctx := context.Background()
	sid := "123"
	exist, err := s.SessionExist(ctx, sid)
	if err != nil {
		t.Errorf("SessionExist error: %v", err)
	}
	if !exist {
		t.Errorf("SessionExist failed")
	}
}
