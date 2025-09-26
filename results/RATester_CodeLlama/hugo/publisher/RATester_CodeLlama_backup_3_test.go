package publisher

import (
	"fmt"
	"testing"
)

func TestBackup_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &htmlElementsCollectorWriter{}
	l.pos = 10
	l.width = 10
	l.r = 'a'
	l.backup()
	if l.pos != 0 {
		t.Errorf("l.pos = %d, want 0", l.pos)
	}
	if l.r != 'a' {
		t.Errorf("l.r = %c, want 'a'", l.r)
	}
}
