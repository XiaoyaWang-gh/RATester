package identity

import (
	"fmt"
	"testing"
)

func TestPutSearchID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sid := &searchID{}
	putSearchID(sid)
	if sid.id != nil {
		t.Errorf("sid.id should be nil")
	}
	if sid.isDp != false {
		t.Errorf("sid.isDp should be false")
	}
	if sid.isPeq != false {
		t.Errorf("sid.isPeq should be false")
	}
	if sid.hasEqer != false {
		t.Errorf("sid.hasEqer should be false")
	}
	if sid.maxDepth != 0 {
		t.Errorf("sid.maxDepth should be 0")
	}
	if sid.dp != nil {
		t.Errorf("sid.dp should be nil")
	}
	if sid.peq != nil {
		t.Errorf("sid.peq should be nil")
	}
	if sid.eqer != nil {
		t.Errorf("sid.eqer should be nil")
	}
	if len(sid.seen) != 0 {
		t.Errorf("len(sid.seen) should be 0")
	}
}
