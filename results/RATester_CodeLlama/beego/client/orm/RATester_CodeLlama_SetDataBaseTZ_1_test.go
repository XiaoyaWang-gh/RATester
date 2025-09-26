package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDataBaseTZ_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	aliasName := "test"
	tz := time.UTC
	err := SetDataBaseTZ(aliasName, tz)
	if err != nil {
		t.Errorf("SetDataBaseTZ() error = %v", err)
		return
	}
	if al, ok := dataBaseCache.get(aliasName); ok {
		if al.TZ != tz {
			t.Errorf("SetDataBaseTZ() al.TZ = %v, want %v", al.TZ, tz)
		}
	} else {
		t.Errorf("SetDataBaseTZ() dataBaseCache.get() error")
	}
}
