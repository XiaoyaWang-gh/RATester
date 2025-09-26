package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestIsAllDatesZero_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := Dates{
		Date:        time.Time{},
		Lastmod:     time.Time{},
		PublishDate: time.Time{},
		ExpiryDate:  time.Time{},
	}
	if !d.IsAllDatesZero() {
		t.Errorf("Expected all dates to be zero")
	}
}
