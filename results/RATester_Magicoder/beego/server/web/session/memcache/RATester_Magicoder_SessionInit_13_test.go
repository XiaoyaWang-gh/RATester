package memcache

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSessionInit_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &MemProvider{}
	maxlifetime := int64(3600)
	savePath := "127.0.0.1:11211;127.0.0.1:11212"
	err := rp.SessionInit(ctx, maxlifetime, savePath)
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}
	if rp.maxlifetime != maxlifetime {
		t.Errorf("Expected maxlifetime to be %d, but got %d", maxlifetime, rp.maxlifetime)
	}
	if !reflect.DeepEqual(rp.conninfo, strings.Split(savePath, ";")) {
		t.Errorf("Expected conninfo to be %v, but got %v", strings.Split(savePath, ";"), rp.conninfo)
	}
}
