package resources

import (
	"fmt"
	"testing"
)

func TestSetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{}
	name := "test"
	r.setName(name)
	if r.name != name {
		t.Errorf("r.name = %s, want %s", r.name, name)
	}
	if !r.changed {
		t.Errorf("r.changed = %v, want %v", r.changed, true)
	}
}
