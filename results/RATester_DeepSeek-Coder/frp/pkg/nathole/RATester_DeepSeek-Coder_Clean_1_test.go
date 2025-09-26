package nathole

import (
	"fmt"
	"testing"
	"time"
)

func TestClean_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Analyzer{
		records: map[string]*MakeHoleRecords{
			"192.168.1.1:192.168.1.2": {
				LastUpdateTime: time.Now().Add(-2 * time.Hour),
			},
			"192.168.1.3:192.168.1.4": {
				LastUpdateTime: time.Now().Add(-3 * time.Hour),
			},
		},
		dataReserveDuration: 1 * time.Hour,
	}

	count, total := a.Clean()
	if count != 1 || total != 2 {
		t.Errorf("expected count 1 and total 2, got count %d and total %d", count, total)
	}

	if _, ok := a.records["192.168.1.1:192.168.1.2"]; ok {
		t.Errorf("expected record to be deleted")
	}

	if _, ok := a.records["192.168.1.3:192.168.1.4"]; !ok {
		t.Errorf("expected record not to be deleted")
	}
}
