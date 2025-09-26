package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		maxlifetime: 10,
		savePath:    "./tmp",
	}
	ctx := context.Background()
	oldsid := "1234567890"
	sid := "12345678901234567890123456789012"
	_, err := fp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Fatal(err)
	}
}
