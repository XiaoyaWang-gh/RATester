package identity

import (
	"fmt"
	"testing"
)

func TestgetSearchID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := getSearchID()

	if id == nil {
		t.Error("Expected a searchID, but got nil")
	}

	if id.id != nil {
		t.Error("Expected id to be nil, but got", id.id)
	}

	if id.isDp != false {
		t.Error("Expected isDp to be false, but got", id.isDp)
	}

	if id.isPeq != false {
		t.Error("Expected isPeq to be false, but got", id.isPeq)
	}

	if id.hasEqer != false {
		t.Error("Expected hasEqer to be false, but got", id.hasEqer)
	}

	if id.maxDepth != 0 {
		t.Error("Expected maxDepth to be 0, but got", id.maxDepth)
	}

	if id.seen != nil {
		t.Error("Expected seen to be nil, but got", id.seen)
	}

	if id.dp != nil {
		t.Error("Expected dp to be nil, but got", id.dp)
	}

	if id.peq != nil {
		t.Error("Expected peq to be nil, but got", id.peq)
	}

	if id.eqer != nil {
		t.Error("Expected eqer to be nil, but got", id.eqer)
	}
}
