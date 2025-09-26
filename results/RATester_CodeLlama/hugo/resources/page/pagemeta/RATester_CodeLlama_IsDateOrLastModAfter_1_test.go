package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestIsDateOrLastModAfter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := Dates{
		Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		Lastmod:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		PublishDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		ExpiryDate:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	in := Dates{
		Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		Lastmod:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		PublishDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		ExpiryDate:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	if !d.IsDateOrLastModAfter(in) {
		t.Errorf("d.IsDateOrLastModAfter(in) = false, want true")
	}
}
