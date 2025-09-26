package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	lp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp",
		Db:          0,
	}
	lp.SessionGC(ctx)
}
