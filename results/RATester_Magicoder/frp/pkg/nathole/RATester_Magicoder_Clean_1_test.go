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
		records:             make(map[string]*MakeHoleRecords),
		dataReserveDuration: time.Hour,
	}

	// Add some records
	a.records["1.1.1.1:2.2.2.2"] = &MakeHoleRecords{
		LastUpdateTime: time.Now().Add(-2 * time.Hour),
	}
	a.records["2.2.2.2:3.3.3.3"] = &MakeHoleRecords{
		LastUpdateTime: time.Now().Add(-3 * time.Hour),
	}

	count, total := a.Clean()

	if count != 1 || total != 2 {
		t.Errorf("Expected count: 1, total: 2, but got count: %d, total: %d", count, total)
	}

	if _, ok := a.records["1.1.1.1:2.2.2.2"]; ok {
		t.Error("Record 1.1.1.1:2.2.2.2 should have been deleted")
	}

	if _, ok := a.records["2.2.2.2:3.3.3.3"]; !ok {
		t.Error("Record 2.2.2.2:3.3.3.3 should not have been deleted")
	}
}
