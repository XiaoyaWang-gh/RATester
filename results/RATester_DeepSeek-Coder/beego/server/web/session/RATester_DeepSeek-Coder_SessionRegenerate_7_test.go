package session

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestSessionRegenerate_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 1800,
		savePath:    "/tmp",
	}

	ctx := context.Background()

	oldsid := "oldsid"
	sid := "newsid"

	_, err := fp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = fp.SessionRegenerate(ctx, oldsid, sid)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	os.Remove(filepath.Join(fp.savePath, string(sid[0]), string(sid[1]), sid))
}
