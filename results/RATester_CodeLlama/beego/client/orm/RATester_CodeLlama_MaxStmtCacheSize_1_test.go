package orm

import (
	"fmt"
	"testing"
)

func TestMaxStmtCacheSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := 100
	opt := MaxStmtCacheSize(v)
	al := &alias{}
	opt(al)
	if al.StmtCacheSize != v {
		t.Errorf("al.StmtCacheSize = %d, want %d", al.StmtCacheSize, v)
	}
}
